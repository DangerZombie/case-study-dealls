package user_repository

import (
	"github.com/DangerZombie/case-study-dealls/model/entity"
	"github.com/DangerZombie/case-study-dealls/model/parameter"
	"gorm.io/gorm"
)

func (r *userRepo) ResetSwipeCount(db *gorm.DB, input parameter.ResetSwipeCountInput) (output parameter.ResetSwipeCountOutput, err error) {
	var user entity.User

	err = db.
		Model(&user).
		Where("status = ?", input.Status).
		Update("swipe_count", input.SwipeCount).
		Error

	if err != nil {
		return output, err
	}

	output.Message = "Success"

	return
}
