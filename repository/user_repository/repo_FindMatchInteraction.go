package user_repository

import (
	"errors"

	"github.com/DangerZombie/case-study-dealls/helper/static"
	"github.com/DangerZombie/case-study-dealls/model/entity"
	"github.com/DangerZombie/case-study-dealls/model/parameter"
	"gorm.io/gorm"
)

func (r *userRepo) FindMatchInteraction(db *gorm.DB, input parameter.FindMatchInteractionInput) (output parameter.FindMatchInteractionOutput, err error) {
	var userInteraction entity.UserInteraction

	err = db.Where("user_id1 = ? AND user_id2 = ? AND interaction_type = ? ", input.UserId2, input.UserId1, static.InteractionTypeLike).
		First(&userInteraction).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return output, nil
		}

		return output, err
	}

	output = parameter.FindMatchInteractionOutput(userInteraction)

	return
}
