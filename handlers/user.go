package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/michelaquino/golang_api_skeleton/context"
	apiMiddleware "github.com/michelaquino/golang_api_skeleton/middleware"
	"github.com/michelaquino/golang_api_skeleton/models"
	"github.com/michelaquino/golang_api_skeleton/repository"
)

// UserHandler is a struct that
type UserHandler struct {
	userRepository repository.UserRepository
}

// NewUserHandler return a new pointer of user struct
func NewUserHandler(userRepository repository.UserRepository) *UserHandler {
	return &UserHandler{
		userRepository: userRepository,
	}
}

// CreateUser is a hendler that create a new user on database
func (h UserHandler) CreateUser(echoContext echo.Context) error {
	userHandlerLog := context.GetLogger()
	requestLogData := echoContext.Get(apiMiddleware.RequestIDKey).(models.RequestLogData)

	userModel := models.UserModel{
		Name:  "User name",
		Email: "user@mail.com",
	}

	if err := h.userRepository.Insert(requestLogData, userModel); err != nil {
		userHandlerLog.Error("Handlers", "CreateUser", requestLogData.ID, requestLogData.OriginIP, "Create user", "Error", err.Error())
		return echoContext.NoContent(http.StatusInternalServerError)
	}

	userHandlerLog.Info("Handlers", "CreateUser", requestLogData.ID, requestLogData.OriginIP, "Create user", "Success", "")
	return echoContext.NoContent(http.StatusCreated)
}
