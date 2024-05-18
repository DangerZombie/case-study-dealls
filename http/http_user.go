package http

import (
	"github.com/DangerZombie/case-study-dealls/endpoint"
	"github.com/DangerZombie/case-study-dealls/service/user_service"
	"github.com/labstack/echo/v4"
)

func (h *httpImpl) UserHandler(group *echo.Group, s user_service.UserService) {
	group.POST("/login", func(ctx echo.Context) error {
		statusCode, result := endpoint.NewEndpoint(h.authHelper).LoginRequest(ctx, s)
		return ctx.JSON(statusCode, result)
	})

	group.POST("/register", func(ctx echo.Context) error {
		statusCode, result := endpoint.NewEndpoint(h.authHelper).RegisterUserRequest(ctx, s)
		return ctx.JSON(statusCode, result)
	})

	group.GET("/profile/:id", func(ctx echo.Context) error {
		statusCode, result := endpoint.NewEndpoint(h.authHelper).ShowUserProfileRequest(ctx, s)
		return ctx.JSON(statusCode, result)
	})

	group.POST("/swipe", func(ctx echo.Context) error {
		statusCode, result := endpoint.NewEndpoint(h.authHelper).SwipeUserRequest(ctx, s)
		return ctx.JSON(statusCode, result)
	})

	group.POST("/subscribe", func(ctx echo.Context) error {
		statusCode, result := endpoint.NewEndpoint(h.authHelper).SubscriptionRequest(ctx, s)
		return ctx.JSON(statusCode, result)
	})
}
