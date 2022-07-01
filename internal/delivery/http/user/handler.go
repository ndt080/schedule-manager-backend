package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ndt080/schedule-manager-backend/internal/delivery/http/response"
	"github.com/ndt080/schedule-manager-backend/internal/service"
	"net/http"
	"strconv"
)

type Handler struct {
	services *service.Service
}

func NewUserHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

// GetCurrentUser godoc
// @Description  Get current user data
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200      {object}  domain.User
// @Failure      400,500  {object}  error.ServerErrorResponse
// @Router       /user/me [get]
// @Security     AuthorizationKey
func (handler *Handler) GetCurrentUser(context *gin.Context) {
	user, err := handler.services.Authorization.GetAuthorizedUser(context)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, response.NewServerInternalError(err.Error()))
		return
	}
	user.PasswordHash = ""
	context.JSON(http.StatusOK, user)
}

// GetUsers godoc
// @Description  Get users data
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        input    body      UsersDataRequest  true  "User ids"
// @Success      200      {object}  UsersDataResponse
// @Failure      400,500  {object}  error.ServerErrorResponse
// @Router       /users [get]
// @Security     AuthorizationKey
func (handler *Handler) GetUsers(context *gin.Context) {
	json := UsersDataRequest{}
	if err := context.ShouldBindJSON(&json); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	users, err := handler.services.User.GetUsersById(json.IDs)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"users": users,
	})
}

// GetUserById godoc
// @Description  Get user data by id
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id       path      string  true  "User id"
// @Success      200      {object}  domain.User
// @Failure      400,500  {object}  error.ServerErrorResponse
// @Router       /user/{id} [get]
// @Security     AuthorizationKey
func (handler *Handler) GetUserById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	user, err := handler.services.User.GetUserById(id)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, response.NewServerBadRequestError(err.Error()))
		return
	}

	context.JSON(http.StatusOK, user)
}
