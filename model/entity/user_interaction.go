package entity

import "github.com/DangerZombie/case-study-dealls/model/base"

type UserInteraction struct {
	base.BaseModel

	UserId1         string `gorm:"type:varchar" json:"-"`
	UserId2         string `gorm:"type:varchar" json:"-"`
	InteractionType string `gorm:"type:varchar" json:"-"`

	User1 User `gorm:"foreignKey:UserId1"`
	User2 User `gorm:"foreignKey:UserId2"`
}
