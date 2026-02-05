package main

import (
	"homework_submit/dao"
	"homework_submit/router"
)

func composer() {
	err := dao.InitDb()
	if err != nil {
		panic(err)
	}
	service := router.Router()
	err = service.Start(":8080")
	if err != nil {
		panic(err)
	}
}
