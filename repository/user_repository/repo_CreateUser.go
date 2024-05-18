package user_repository

import (
	"github.com/DangerZombie/case-study-dealls/model/parameter"
	"gorm.io/gorm"
)

func (r *userRepo) CreateUser(db *gorm.DB, input parameter.CreateUserInput) (output parameter.CreateUserOutput, err error) {
	err = db.Create(&input.User).Error
	if err != nil {
		return output, err
	}

	output.Id = input.Id

	return
}
