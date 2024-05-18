package user_service

import (
	"github.com/DangerZombie/case-study-dealls/model/request"
	"github.com/DangerZombie/case-study-dealls/model/response"
)

type UserService interface {
	BuySubscription(req request.BuySubscriptionRequestBody) (res response.BuySubscriptionResponse, code int, err error)
	CreateUserInteraction(req request.CreateUserInteractionRequest) (res response.CreateUserInteractionResponse, code int, err error)
	GetUserToSwipe(req request.GetUserToSwipeRequest) (res response.GetUserToSwipeResponse, code int, err error)
	Login(req request.LoginRequestBody) (res response.LoginResponse, code int, err error)
	RegisterUser(req request.RegisterUserRequestBody) (res response.RegisterUserResponse, code int, err error)
	ResetSwipeCount(req request.ResetSwipeCountRequest) (res response.ResetSwipeCountResponse, err error)
}
