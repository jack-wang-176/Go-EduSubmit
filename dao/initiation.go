package dao

import (
	"homework_submit/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var sqlStr = "root:123456@@tcp(127.0.0.1:3306)/winter_project?charset=utf8mb4&parseTime=True&loc=local"

func InitDb() error {
	var err error
	DB, err = gorm.Open(mysql.Open(sqlStr), &gorm.Config{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&model.Homework{}, &model.Submission{}, &model.User{})
	if err != nil {
		return err
	}
	return nil
}
