package user_repository

import (
	"github.com/DangerZombie/case-study-dealls/helper/static"
	"github.com/DangerZombie/case-study-dealls/model/parameter"
	"gorm.io/gorm"
)

func (r *userRepo) CreateUserInteraction(db *gorm.DB, input parameter.CreateUserInteractionInput) (output parameter.CreateUserInteractionOutput, err error) {
	err = db.Create(&input.UserInteraction).Error
	if err != nil {
		return output, err
	}

	output.Message = static.MessageDislike
	if input.InteractionType == static.InteractionTypeLike {
		output.Message = static.MessageLike
	}

	return
}
