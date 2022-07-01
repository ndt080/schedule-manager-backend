package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/ndt080/schedule-manager-backend/internal/delivery/http/response"
	"github.com/ndt080/schedule-manager-backend/internal/domain"
	"github.com/ndt080/schedule-manager-backend/internal/service"
	manager "github.com/ndt080/schedule-manager-backend/pkg/auth"
	"net/http"
)

type Handler struct {
	services     *service.Service
	tokenManager manager.TokenManager
}

func NewAuthHandler(services *service.Service, tokenManager manager.TokenManager) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
	}
}

// SignIn godoc
// @Description  Authorization in the account
// @Tags         Authorization
// @Accept       json
// @Produce      json
// @Param        input    body      Credential  true  "Enter user credential"
// @Success      200      {object}  UserResponse
// @Failure      400,500  {object}  error.ServerErrorResponse
// @Router       /auth/sign-in [post]
func (handler *Handler) SignIn(context *gin.Context) {
	json := Credential{}
	if err := context.ShouldBindJSON(&json); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	user, err := handler.services.User.GetUserByEmail(json.Email)
	if err != nil || user == nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, response.NewServerInvalidCredentialsError())
		return
	}

	if isValid := handler.tokenManager.CheckPasswordHash(json.Password, user.PasswordHash); !isValid {
		context.AbortWithStatusJSON(http.StatusInternalServerError, response.NewServerInvalidCredentialsError())
		return
	}

	tokens, err := handler.services.Authorization.CreateTokens(user.ID)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, response.NewServerInternalError(err.Error()))
		return
	}

	user.PasswordHash = ""
	context.JSON(http.StatusOK, map[string]interface{}{
		"user":   user,
		"tokens": tokens,
	})
}

// SignUp godoc
// @Description  Creating an account
// @Tags         Authorization
// @Accept       json
// @Produce      json
// @Param        input    body      UserRequest  true  "Enter user credential"
// @Success      200      {object}  domain.User
// @Failure      400,500  {object}  error.ServerErrorResponse
// @Router       /auth/sign-up [post]
func (handler *Handler) SignUp(context *gin.Context) {
	json := UserRequest{}
	if err := context.ShouldBindJSON(&json); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	if isExists, _ := handler.services.User.CheckExistsUser(json.Email); !isExists {
		context.AbortWithStatusJSON(http.StatusInternalServerError, response.NewServerCredentialsExistsError())
		return
	}

	user, err := handler.services.User.CreateUser(&domain.User{
		Username:     json.Username,
		Email:        json.Email,
		PasswordHash: json.Password,
	})
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, response.NewServerInternalError(err.Error()))
		return
	}

	if _, err := handler.services.Authorization.SendVerificationEmail(user, context); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, response.NewServerInternalError(err.Error()))
		return
	}

	user.PasswordHash = ""
	context.JSON(http.StatusOK, user)
}

// ConfirmEmailAgain godoc
// @Description  Confirm email again
// @Tags         Authorization
// @Accept       json
// @Produce      json
// @Param        email    query     string  false  "User email"
// @Success      200      {object}  domain.User
// @Failure      404,500  {object}  error.ServerErrorResponse
// @Router       /auth/confirm-email-again [post]
func (handler *Handler) ConfirmEmailAgain(context *gin.Context) {
	email, ok := context.GetQuery("email")
	if !ok {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError("Invalid verification token"))
		return
	}

	user, err := handler.services.User.GetUserByEmail(email)
	if user == nil || err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, response.NewServerInvalidCredentialsError())
		return
	}

	if user.IsVerified {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError("The account has already been verified"))
		return
	}

	if _, err := handler.services.Authorization.SendVerificationEmail(user, context); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, response.NewServerInternalError(err.Error()))
		return
	}

	user.PasswordHash = ""
	context.JSON(http.StatusOK, user)
}

// RefreshToken godoc
// @Description  Refresh token
// @Tags         Authorization
// @Accept       json
// @Produce      json
// @Param        input        body      RefreshTokenRequest  true  "Enter refresh token"
// @Success      200          {object}  RefreshTokenResponse
// @Failure      400,401,500  {object}  error.ServerErrorResponse
// @Router       /auth/refresh-token [post]
func (handler *Handler) RefreshToken(context *gin.Context) {
	json := RefreshTokenRequest{}
	if err := context.ShouldBindJSON(&json); err != nil || json.RefreshToken == "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	claims, err := handler.tokenManager.ParseToken(json.RefreshToken)
	if err != nil || claims.Valid() != nil || claims.TokenType != "refresh" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, response.NewServerInvalidRefreshTokenError())
		return
	}

	tokens, err := handler.services.Authorization.CreateTokens(claims.UserId)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, response.NewServerInternalError(err.Error()))
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"tokens": tokens,
	})
}

// CheckAuthStatus godoc
// @Description  Check authorization status
// @Tags         Authorization
// @Accept       json
// @Produce      json
// @Success      200          {object}  success.ServerSuccessResponse
// @Failure      400,401,500  {object}  error.ServerErrorResponse
// @Router       /auth/status [get]
// @Security     AuthorizationKey
func (handler *Handler) CheckAuthStatus(context *gin.Context) {
	context.JSON(http.StatusOK, response.NewServerSuccessResponse("Ok"))
}

// VerifyEmail godoc
// @Description  Verify email address
// @Tags         Authorization
// @Accept       json
// @Produce      json
// @Param        token        query     string  false  "Verification token"
// @Failure      400,401,500  {object}  error.ServerErrorResponse
// @Router       /auth/verify-email [get]
func (handler *Handler) VerifyEmail(context *gin.Context) {
	token, ok := context.GetQuery("token")
	if !ok {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError("Invalid verification token"))
		return
	}

	claims, err := handler.tokenManager.ParseToken(token)
	if err != nil || claims.Valid() != nil || claims.TokenType != "verification" {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError("Invalid verification token"))
		return
	}

	if err := handler.services.User.ConfirmUserVerification(claims.UserId); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, response.NewServerInternalError(err.Error()))
		return
	}

	context.Redirect(http.StatusMovedPermanently, "https://schedule-manager.darkguin.com")
}
