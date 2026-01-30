package dao

import "homework_submit/model"

type UserDao struct{}

var User = new(UserDao)

func (d *UserDao) CreateUser(u *model.User) error {
	return DB.Create(u).Error
}
func (d *UserDao) DeleteUser(u *model.User) error {
	return DB.Delete(u).Error
}

//TODO 用户登录和获取用户信息,密码加密和刷新token
