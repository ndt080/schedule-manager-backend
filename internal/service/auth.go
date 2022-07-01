package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ndt080/schedule-manager-backend/internal/delivery/http/response"
	"github.com/ndt080/schedule-manager-backend/internal/delivery/smtp"
	"github.com/ndt080/schedule-manager-backend/internal/domain"
	"github.com/ndt080/schedule-manager-backend/internal/repository"
	"github.com/ndt080/schedule-manager-backend/pkg/auth"
	"net/http"
	"strings"
	"time"
)

const authorizationHeader = "Authorization"

type AuthService struct {
	tokenManager   auth.TokenManager
	smtpService    *smtp.SmtpService
	userRepository repository.User
}

func NewAuthService(tokenManager auth.TokenManager, smtpService *smtp.SmtpService, userRepository repository.User) *AuthService {
	return &AuthService{
		tokenManager:   tokenManager,
		smtpService:    smtpService,
		userRepository: userRepository,
	}
}

func (service *AuthService) CreateTokens(userId int64) (*domain.Tokens, error) {
	accessToken, err := service.tokenManager.NewAccessToken(userId)
	if err != nil {
		return nil, err
	}

	refreshToken, err := service.tokenManager.NewRefreshToken(userId)
	if err != nil {
		return nil, err
	}

	return &domain.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (service *AuthService) SendVerificationEmail(user *domain.User, context *gin.Context) (bool, error) {
	token, err := service.tokenManager.NewVerificationToken(user.ID, 24*time.Hour)
	if err != nil {
		return false, err
	}

	host := context.Request.Host
	templateData := struct{ URL string }{
		URL: fmt.Sprintf("https://%s/auth/verify-email?token=%s", host, token),
	}

	body, err := service.smtpService.ParseTemplate("confirm.html", templateData)
	if err != nil {
		return false, err
	}

	return service.smtpService.SendEmail(smtp.SmtpRequest{
		To:      []string{user.Email},
		Subject: "Confirm email",
		Body:    body,
	})
}

func (service *AuthService) GetAuthorizedUser(context *gin.Context) (*domain.User, error) {
	userId, _ := context.Get("userId")
	return service.userRepository.GetUserById(userId.(int64))
}

func (service *AuthService) IdentifyUser(context *gin.Context) {
	header := context.GetHeader(authorizationHeader)
	if header == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, response.NewServerInvalidAccessTokenError())
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		context.AbortWithStatusJSON(http.StatusUnauthorized, response.NewServerInvalidAccessTokenError())
		return
	}

	claims, err := service.tokenManager.ParseToken(headerParts[1])
	if err != nil || claims.TokenType != "jwt" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, response.NewServerInvalidAccessTokenError())
		return
	}

	context.Set("userId", claims.UserId)
}
