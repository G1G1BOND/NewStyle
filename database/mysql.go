package database

import (
	"fmt"
	"go-v1/model"
	"go-v1/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"v1/model"
)

var userDB *gorm.DB
var momentDB *gorm.DB

//type UserDao struct {
//	db *gorm.DB
//}

func InitDB() {
	//host := "localhost"
	//port := "3306"
	//database := "test01"
	//username := "root"
	//password := "111111"
	//charset := "utf8"
	//数据库信息
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
		utils.DbCharset)
	//连接mysql数据库
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("数据库连接错误，错误信息：" + err.Error())
	}
	//迁移
	db.AutoMigrate(&model.User{})
	userDB = db
	db.AutoMigrate(&model.Moment{})
	momentDB = db
}

func GetDB() *gorm.DB {
	return userDB
}
func MomentDB() *gorm.DB {
	return momentDB
}

