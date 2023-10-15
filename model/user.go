package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null;unique"`
	Email    string `gorm:"type:varchar(50);not null;unique"`
	Password string `gorm:"size:255;not null"`
	//Avatar string `gorm:"type:varchar(255);not null"`
	Avatar string `gorm:"default:http://rsa9yybad.hn-bkt.clouddn.com/FmEya-O5gDPkQ55efxU2dSnOZ7UI"`
}

func (*User) TableName() string {
	return "user"
}
