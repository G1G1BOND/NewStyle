package model

import (
	"gorm.io/gorm"
)

type Moment struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	//Email    string `gorm:"type:varchar(50);not null;unique"`
	Content string `gorm:"size:255;not null"`
	Like    string `gorm:"size:10;not null;default:0"`
	//Avatar string `gorm:"type:varchar(255);not null"`
	Avatar  string `gorm:"default:http://rpgydm0qh.hn-bkt.clouddn.com/bc41368afc878f84a970d172dc26d70f.jpg"`
	Picture string `gorm:"size:255;not null"`
}
