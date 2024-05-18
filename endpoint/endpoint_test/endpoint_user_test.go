package endpoint_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DangerZombie/case-study-dealls/endpoint"
	"github.com/DangerZombie/case-study-dealls/helper/auth"
	"github.com/DangerZombie/case-study-dealls/helper/static"
	"github.com/DangerZombie/case-study-dealls/model/parameter"
	"github.com/DangerZombie/case-study-dealls/model/request"
	"github.com/DangerZombie/case-study-dealls/model/response"
	"github.com/DangerZombie/case-study-dealls/service/user_service"
	"github.com/go-faker/faker/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestEndpointUser_Login(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockUserService := user_service.NewMockUserService(mockCtrl)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)

	endpointModule := endpoint.NewEndpoint(
		mockAuthHelper,
	)

	e := echo.New()
	loginRequest := request.LoginRequestBody{
		Username: faker.Username(),
		Password: faker.Name(),
	}

	loginResponse := response.LoginResponse{
		Token: faker.Jwt(),
	}

	t.Run("Should return OK", func(t *testing.T) {
		reqBody, _ := json.Marshal(loginRequest)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/user/login", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockUserService.EXPECT().
			Login(loginRequest).
			Times(1).
			Return(loginResponse, http.StatusOK, nil)

		statusCode, result := endpointModule.LoginRequest(c, mockUserService)

		require.Equal(t, http.StatusOK, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return error", func(t *testing.T) {
		reqBody, _ := json.Marshal(loginRequest)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/user/login", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockUserService.EXPECT().
			Login(loginRequest).
			Times(1).
			Return(loginResponse, http.StatusInternalServerError, errors.New("failed"))

		statusCode, result := endpointModule.LoginRequest(c, mockUserService)

		require.Equal(t, http.StatusInternalServerError, statusCode)
		require.NotEmpty(t, result)
	})
}

func TestEndpointUser_Register(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockUserService := user_service.NewMockUserService(mockCtrl)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)

	endpointModule := endpoint.NewEndpoint(
		mockAuthHelper,
	)

	e := echo.New()
	registerRequest := request.RegisterUserRequestBody{
		Username: faker.Username(),
		Password: faker.Name(),
		Nickname: faker.Name(),
		Gender:   faker.Gender(),
		Age:      30,
		Location: faker.Sentence(),
	}

	registerResponse := response.RegisterUserResponse{
		Message: "Success",
	}

	t.Run("Should return OK", func(t *testing.T) {
		reqBody, _ := json.Marshal(registerRequest)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/user/register", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockUserService.EXPECT().
			RegisterUser(registerRequest).
			Times(1).
			Return(registerResponse, http.StatusOK, nil)

		statusCode, result := endpointModule.RegisterUserRequest(c, mockUserService)

		require.Equal(t, http.StatusOK, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return error", func(t *testing.T) {
		reqBody, _ := json.Marshal(registerRequest)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/user/register", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockUserService.EXPECT().
			RegisterUser(registerRequest).
			Times(1).
			Return(registerResponse, http.StatusInternalServerError, errors.New("failed"))

		statusCode, result := endpointModule.RegisterUserRequest(c, mockUserService)

		require.Equal(t, http.StatusInternalServerError, statusCode)
		require.NotEmpty(t, result)
	})
}

func TestEndpointUser_ShowUserProfile(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockUserService := user_service.NewMockUserService(mockCtrl)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)

	endpointModule := endpoint.NewEndpoint(
		mockAuthHelper,
	)

	e := echo.New()
	id1 := faker.UUIDHyphenated()
	id2 := faker.UUIDHyphenated()
	claims := parameter.JwtClaims{
		Issuer:  id1,
		Subject: faker.Name(),
		User:    faker.Name(),
	}

	showUserProfileRequest := request.GetUserToSwipeRequest{
		Id: id1,
	}

	showUserProfileResponse := response.GetUserToSwipeResponse{
		Id:       id2,
		Nickname: faker.Name(),
		Gender:   faker.Gender(),
		Age:      30,
		Location: faker.Sentence(),
	}

	t.Run("Should return OK", func(t *testing.T) {
		url := fmt.Sprintf("/api/v1/user/profile/%s", id1)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/user/profile/:id")
		c.SetParamNames("id")
		c.SetParamValues(id1)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(claims, nil)

		mockUserService.EXPECT().
			GetUserToSwipe(showUserProfileRequest).
			Times(1).
			Return(showUserProfileResponse, http.StatusOK, nil)

		statusCode, result := endpointModule.ShowUserProfileRequest(c, mockUserService)

		require.Equal(t, http.StatusOK, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return error unauthorized", func(t *testing.T) {
		url := fmt.Sprintf("/api/v1/user/profile/%s", id1)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/teacher/:id")
		c.SetParamNames("id")
		c.SetParamValues(id1)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(parameter.JwtClaims{}, errors.New("failed"))

		statusCode, result := endpointModule.ShowUserProfileRequest(c, mockUserService)

		require.Equal(t, http.StatusUnauthorized, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return error", func(t *testing.T) {
		url := fmt.Sprintf("/api/v1/user/profile/%s", id1)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/user/profile/:id")
		c.SetParamNames("id")
		c.SetParamValues(id1)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(claims, nil)

		mockUserService.EXPECT().
			GetUserToSwipe(showUserProfileRequest).
			Times(1).
			Return(showUserProfileResponse, http.StatusInternalServerError, errors.New("failed"))

		statusCode, result := endpointModule.ShowUserProfileRequest(c, mockUserService)

		require.Equal(t, http.StatusInternalServerError, statusCode)
		require.NotEmpty(t, result)
	})
}

func TestEndpointUser_Swipe(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockUserService := user_service.NewMockUserService(mockCtrl)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)

	endpointModule := endpoint.NewEndpoint(
		mockAuthHelper,
	)

	e := echo.New()
	id1 := faker.UUIDHyphenated()
	id2 := faker.UUIDHyphenated()
	claims := parameter.JwtClaims{
		Issuer:  id1,
		Subject: faker.Name(),
		User:    faker.Name(),
	}

	swipeRequest := request.CreateUserInteractionRequest{
		UserId1:         id1,
		UserId2:         id2,
		InteractionType: static.InteractionTypeLike,
	}

	swipeResponse := response.CreateUserInteractionResponse{
		Message: static.MessageLike,
	}

	t.Run("Should return OK", func(t *testing.T) {
		reqBody, _ := json.Marshal(swipeRequest)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/user/swipe", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(claims, nil)

		mockUserService.EXPECT().
			CreateUserInteraction(swipeRequest).
			Times(1).
			Return(swipeResponse, http.StatusOK, nil)

		statusCode, result := endpointModule.SwipeUserRequest(c, mockUserService)

		require.Equal(t, http.StatusOK, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return error unauthorized", func(t *testing.T) {
		reqBody, _ := json.Marshal(swipeRequest)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/user/swipe", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(parameter.JwtClaims{}, errors.New("failed"))

		statusCode, result := endpointModule.SwipeUserRequest(c, mockUserService)

		require.Equal(t, http.StatusUnauthorized, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return error internal server error", func(t *testing.T) {
		reqBody, _ := json.Marshal(swipeRequest)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/user/swipe", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(claims, nil)

		mockUserService.EXPECT().
			CreateUserInteraction(swipeRequest).
			Times(1).
			Return(swipeResponse, http.StatusInternalServerError, errors.New("failed"))

		statusCode, result := endpointModule.SwipeUserRequest(c, mockUserService)

		require.Equal(t, http.StatusInternalServerError, statusCode)
		require.NotEmpty(t, result)
	})
}

func TestEndpointUser_Subscribe(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockUserService := user_service.NewMockUserService(mockCtrl)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)

	endpointModule := endpoint.NewEndpoint(
		mockAuthHelper,
	)

	e := echo.New()
	id := faker.UUIDHyphenated()
	claims := parameter.JwtClaims{
		Issuer:  id,
		Subject: faker.Name(),
		User:    faker.Name(),
	}

	subscribeRequest := request.BuySubscriptionRequestBody{
		Id: id,
	}

	subscribeResponse := response.BuySubscriptionResponse{
		Message: "Subcsription upgrade to Premium User",
	}

	t.Run("Should return OK", func(t *testing.T) {
		reqBody, _ := json.Marshal(subscribeRequest)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/user/subscribe", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(claims, nil)

		mockUserService.EXPECT().
			BuySubscription(subscribeRequest).
			Times(1).
			Return(subscribeResponse, http.StatusOK, nil)

		statusCode, result := endpointModule.SubscriptionRequest(c, mockUserService)

		require.Equal(t, http.StatusOK, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return error unauthorized", func(t *testing.T) {
		reqBody, _ := json.Marshal(subscribeRequest)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/user/subscribe", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(parameter.JwtClaims{}, errors.New("failed"))

		statusCode, result := endpointModule.SubscriptionRequest(c, mockUserService)

		require.Equal(t, http.StatusUnauthorized, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return error internal server error", func(t *testing.T) {
		reqBody, _ := json.Marshal(subscribeRequest)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/user/subscribe", strings.NewReader(string(reqBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request().Header).
			Times(1).
			Return(claims, nil)

		mockUserService.EXPECT().
			BuySubscription(subscribeRequest).
			Times(1).
			Return(subscribeResponse, http.StatusInternalServerError, errors.New("failed"))

		statusCode, result := endpointModule.SubscriptionRequest(c, mockUserService)

		require.Equal(t, http.StatusInternalServerError, statusCode)
		require.NotEmpty(t, result)
	})
}
