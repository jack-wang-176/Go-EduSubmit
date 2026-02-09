package dao

import (
	"homework_submit/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var sqlStr = "root:123456@tcp(maple-mysql:3306)/winter_project?charset=utf8mb4&parseTime=True&loc=Local"

func InitDb() error {
	var err error
	DB, err = gorm.Open(mysql.Open(sqlStr), &gorm.Config{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&model.Homework{}, &model.Submission{}, &model.User{}, &model.RefreshToken{})
	if err != nil {
		return err
	}
	return nil
}
