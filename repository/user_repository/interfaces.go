package user_repository

import (
	"github.com/DangerZombie/case-study-dealls/model/parameter"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUserInteraction(db *gorm.DB, input parameter.CreateUserInteractionInput) (output parameter.CreateUserInteractionOutput, err error)
	CreateUser(db *gorm.DB, input parameter.CreateUserInput) (output parameter.CreateUserOutput, err error)
	CreateMatchUser(db *gorm.DB, input parameter.CreateMatchUserInput) (output parameter.CreateMatchUserOutput, err error)
	FindMatchInteraction(db *gorm.DB, input parameter.FindMatchInteractionInput) (output parameter.FindMatchInteractionOutput, err error)
	FindUserById(db *gorm.DB, input parameter.FindUserByIdInput) (output parameter.FindUserByIdOutput, err error)
	FindUserByUsernameAndPassword(db *gorm.DB, input parameter.FindUserByUsernameAndPasswordInput) (output parameter.FindUserByUsernameAndPasswordOutput, err error)
	FindUserToSwipe(db *gorm.DB, input parameter.FindUserToSwipeInput) (output parameter.FindUserToSwipeOutput, err error)
	ResetSwipeCount(db *gorm.DB, input parameter.ResetSwipeCountInput) (output parameter.ResetSwipeCountOutput, err error)
	UpdateSubscription(db *gorm.DB, input parameter.UpdateSubscriptionInput) (output parameter.UpdateSubscriptionOutput, err error)
	UpdateSwipeCount(db *gorm.DB, input parameter.UpdateSwipeCountInput) (output parameter.UpdateSwipeCountOutput, err error)
}
