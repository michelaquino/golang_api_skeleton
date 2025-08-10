package handlers

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/michelaquino/golang_api_skeleton/src/models"
	"github.com/michelaquino/golang_api_skeleton/src/repository"
)

// UserHandler is a struct that stores an userRepository.
type UserHandler struct {
	userRepository repository.UserRepository
}

// NewUserHandler returns a new pointer of user's struct.
func NewUserHandler(userRepository repository.UserRepository) *UserHandler {
	return &UserHandler{
		userRepository: userRepository,
	}
}

// CreateUser is a handler that creates a new user into database.
func (h UserHandler) CreateUser(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()
	userModel := models.UserModel{}
	if err := echoContext.Bind(&userModel); err != nil {
		slog.ErrorContext(ctx, err.Error())
		return echoContext.NoContent(http.StatusBadRequest)
	}

	if err := h.userRepository.Insert(echoContext.Request().Context(), userModel); err != nil {
		slog.ErrorContext(ctx, err.Error())
		return echoContext.NoContent(http.StatusInternalServerError)
	}

	slog.InfoContext(ctx, "user created with success")
	return echoContext.NoContent(http.StatusCreated)
}
