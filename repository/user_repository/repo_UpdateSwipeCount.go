package user_repository

import (
	"github.com/DangerZombie/case-study-dealls/model/entity"
	"github.com/DangerZombie/case-study-dealls/model/parameter"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *userRepo) UpdateSwipeCount(db *gorm.DB, input parameter.UpdateSwipeCountInput) (output parameter.UpdateSwipeCountOutput, err error) {
	var user entity.User

	err = db.
		Model(&user).
		Where("id = ?", input.Id).
		Clauses(clause.Returning{}).
		Update("swipe_count", input.SwipeCount).
		Error

	if err != nil {
		return output, err
	}

	output = parameter.UpdateSwipeCountOutput{
		User: user,
	}

	return output, nil
}
