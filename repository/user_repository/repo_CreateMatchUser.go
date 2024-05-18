package user_repository

import (
	"github.com/DangerZombie/case-study-dealls/model/parameter"
	"gorm.io/gorm"
)

func (r *userRepo) CreateMatchUser(db *gorm.DB, input parameter.CreateMatchUserInput) (output parameter.CreateMatchUserOutput, err error) {
	err = db.Create(&input.Match).Error
	if err != nil {
		return output, err
	}

	output.Id = input.Id

	return
}
