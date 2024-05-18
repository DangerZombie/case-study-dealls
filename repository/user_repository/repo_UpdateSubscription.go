package user_repository

import (
	"github.com/DangerZombie/case-study-dealls/helper/static"
	"github.com/DangerZombie/case-study-dealls/model/entity"
	"github.com/DangerZombie/case-study-dealls/model/parameter"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *userRepo) UpdateSubscription(db *gorm.DB, input parameter.UpdateSubscriptionInput) (output parameter.UpdateSubscriptionOutput, err error) {
	var user entity.User

	err = db.
		Model(&user).
		Where("id = ?", input.Id).
		Clauses(clause.Returning{}).
		Updates(map[string]interface{}{
			"status":      static.UserPremium,
			"swipe_count": -1,
			"verified":    true,
		}).
		Error

	if err != nil {
		return output, err
	}

	output = parameter.UpdateSubscriptionOutput{
		Message: "Subcsription upgrade to Premium User",
	}

	return
}
