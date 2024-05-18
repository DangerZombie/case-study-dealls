package http

import (
	"github.com/DangerZombie/case-study-dealls/service/user_service"
	"github.com/labstack/echo/v4"
)

type Http interface {
	// API handler
	UserHandler(group *echo.Group, s user_service.UserService)
}
