package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	transport "github.com/DangerZombie/case-study-dealls/http"

	"github.com/DangerZombie/case-study-dealls/helper/auth"
	"github.com/DangerZombie/case-study-dealls/service/user_service"
	"github.com/labstack/echo/v4"
	"go.uber.org/mock/gomock"
)

func TestHttpUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)
	mockUserService := user_service.NewMockUserService(mockCtrl)

	httpModule := transport.NewHttp(
		mockAuthHelper,
	)

	e := echo.New()
	g := e.Group("/api/v1/user")

	t.Run("Should invoke login", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/login", nil)
		rec := httptest.NewRecorder()
		_ = e.NewContext(req, rec)

		httpModule.UserHandler(g, mockUserService)
	})

	t.Run("Should invoke register", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/register", nil)
		rec := httptest.NewRecorder()
		_ = e.NewContext(req, rec)

		httpModule.UserHandler(g, mockUserService)
	})

	t.Run("Should invoke profile to swipe", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/profile/:id", nil)
		rec := httptest.NewRecorder()
		_ = e.NewContext(req, rec)

		httpModule.UserHandler(g, mockUserService)
	})

	t.Run("Should invoke swipe", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/swipe", nil)
		rec := httptest.NewRecorder()
		_ = e.NewContext(req, rec)

		httpModule.UserHandler(g, mockUserService)
	})
}
