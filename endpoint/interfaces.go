package endpoint

import (
	"github.com/DangerZombie/case-study-dealls/service/user_service"
	"github.com/labstack/echo/v4"
)

type Endpoint interface {
	// Endpoint User
	LoginRequest(ctx echo.Context, s user_service.UserService) (int, interface{})
	RegisterUserRequest(ctx echo.Context, s user_service.UserService) (int, interface{})
	ShowUserProfileRequest(ctx echo.Context, s user_service.UserService) (int, interface{})
	SubscriptionRequest(ctx echo.Context, s user_service.UserService) (int, interface{})
	SwipeUserRequest(ctx echo.Context, s user_service.UserService) (int, interface{})
}
