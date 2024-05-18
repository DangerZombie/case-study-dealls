package parameter

import "github.com/DangerZombie/case-study-dealls/model/entity"

type CreateUserInput struct {
	entity.User
}

type CreateUserOutput struct {
	Id string
}

type CreateUserInteractionInput struct {
	entity.UserInteraction
}

type CreateUserInteractionOutput struct {
	Message string
}

type CreateMatchUserInput struct {
	entity.Match
}

type CreateMatchUserOutput struct {
	Id string
}

type FindMatchInteractionInput struct {
	UserId1 string
	UserId2 string
}

type FindMatchInteractionOutput entity.UserInteraction

type FindUserByIdInput struct {
	Id string
}

type FindUserByIdOutput struct {
	Id         string
	Username   string
	Nickname   string
	Status     string
	SwipeCount int
}

type FindUserByUsernameAndPasswordInput struct {
	Username string
	Password string
}

type FindUserByUsernameAndPasswordOutput entity.User

type FindUserToSwipeInput struct {
	Id string
}

type FindUserToSwipeOutput entity.User

type JwtClaims struct {
	Issuer  string
	Subject string
	User    string
}

type ResetSwipeCountInput struct {
	Status     string
	SwipeCount int
}

type ResetSwipeCountOutput struct {
	Message string
}

type UpdateSubscriptionInput struct {
	Id string
}

type UpdateSubscriptionOutput struct {
	Message string
}

type UpdateSwipeCountInput struct {
	Id         string
	SwipeCount int
}

type UpdateSwipeCountOutput struct {
	entity.User
}
