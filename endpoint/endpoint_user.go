package endpoint

import (
	"encoding/json"
	"net/http"

	"github.com/DangerZombie/case-study-dealls/model/base"
	"github.com/DangerZombie/case-study-dealls/model/request"
	"github.com/DangerZombie/case-study-dealls/service/user_service"
	"github.com/labstack/echo/v4"
)

func (e *endpointImpl) LoginRequest(ctx echo.Context, s user_service.UserService) (int, interface{}) {
	req := request.LoginRequestBody{}
	_ = json.NewDecoder(ctx.Request().Body).Decode(&req)
	result, code, err := s.Login(req)

	var wrap interface{}
	if err != nil {
		wrap = base.SetHttpResponse(nil, err.Error())
	} else {
		wrap = base.SetHttpResponse(result, "")
	}

	return code, wrap
}

func (e *endpointImpl) RegisterUserRequest(ctx echo.Context, s user_service.UserService) (int, interface{}) {
	req := request.RegisterUserRequestBody{}
	_ = json.NewDecoder(ctx.Request().Body).Decode(&req)
	result, code, err := s.RegisterUser(req)

	var wrap interface{}
	if err != nil {
		wrap = base.SetHttpResponse(nil, err.Error())
	} else {
		wrap = base.SetHttpResponse(result, "")
	}

	return code, wrap
}

func (e *endpointImpl) ShowUserProfileRequest(ctx echo.Context, s user_service.UserService) (int, interface{}) {
	// Verify JWT token from the request headers
	_, err := e.authHelper.VerifyJWT(ctx.Request().Header)
	if err != nil {
		wrap := base.SetHttpResponse(nil, err.Error())
		return http.StatusUnauthorized, wrap
	}

	req := request.GetUserToSwipeRequest{
		Id: ctx.Param("id"),
	}

	result, code, err := s.GetUserToSwipe(req)

	var wrap interface{}
	if err != nil {
		wrap = base.SetHttpResponse(nil, err.Error())
	} else {
		wrap = base.SetHttpResponse(result, "")
	}

	return code, wrap
}

func (e *endpointImpl) SwipeUserRequest(ctx echo.Context, s user_service.UserService) (int, interface{}) {
	// Verify JWT token from the request headers
	_, err := e.authHelper.VerifyJWT(ctx.Request().Header)
	if err != nil {
		wrap := base.SetHttpResponse(nil, err.Error())
		return http.StatusUnauthorized, wrap
	}

	req := request.CreateUserInteractionRequest{}
	_ = json.NewDecoder(ctx.Request().Body).Decode(&req)
	result, code, err := s.CreateUserInteraction(req)

	var wrap interface{}
	if err != nil {
		wrap = base.SetHttpResponse(nil, err.Error())
	} else {
		wrap = base.SetHttpResponse(result, "")
	}

	return code, wrap
}

func (e *endpointImpl) SubscriptionRequest(ctx echo.Context, s user_service.UserService) (int, interface{}) {
	// Verify JWT token from the request headers
	_, err := e.authHelper.VerifyJWT(ctx.Request().Header)
	if err != nil {
		wrap := base.SetHttpResponse(nil, err.Error())
		return http.StatusUnauthorized, wrap
	}

	req := request.BuySubscriptionRequestBody{}
	_ = json.NewDecoder(ctx.Request().Body).Decode(&req)
	result, code, err := s.BuySubscription(req)

	var wrap interface{}
	if err != nil {
		wrap = base.SetHttpResponse(nil, err.Error())
	} else {
		wrap = base.SetHttpResponse(result, "")
	}

	return code, wrap
}
