package entity

import "github.com/DangerZombie/case-study-dealls/model/base"

type User struct {
	base.BaseModel

	Username   string `gorm:"type:varchar" json:"-"`
	Password   string `gorm:"type:varchar" json:"-"`
	Status     string `gorm:"type:varchar" json:"status"`
	Nickname   string `gorm:"type:varchar" json:"nickname"`
	Gender     string `gorm:"type:varchar" json:"gender"`
	Age        int    `gorm:"type:int" json:"age"`
	Location   string `gorm:"type:varchar" json:"location"`
	SwipeCount int    `gorm:"type:int" json:"swipe_count"`
	Verified   bool   `gorm:"type:boolean" json:"verified"`

	Interactions []UserInteraction `gorm:"foreignKey:UserId1" json:"user_interaction"`
	Mathces      []Match           `gorm:"foreignKey:UserId1" json:"matches"`
}
