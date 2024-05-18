package user_repository

import (
	"errors"

	"github.com/DangerZombie/case-study-dealls/model/parameter"
	"gorm.io/gorm"
)

func (r userRepo) FindUserToSwipe(db *gorm.DB, input parameter.FindUserToSwipeInput) (output parameter.FindUserToSwipeOutput, err error) {
	query := `
        SELECT u.*
        FROM "user" u
        LEFT JOIN "user_interaction" ui ON u.id = ui.user_id2 AND ui.user_id1 = ?
        WHERE ui.id IS NULL AND u.id != ?
		LIMIT 1
    `
	err = db.Raw(query, input.Id, input.Id).
		Scan(&output).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}

		return output, err
	}

	return
}
