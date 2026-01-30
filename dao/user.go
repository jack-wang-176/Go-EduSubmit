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

func (d *UserDao) GetUserById(id uint) (*model.User, error) {
	var user model.User
	tx := DB.First(&user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (d *UserDao) GetUserByName(name string) (*model.User, error) {
	var user model.User
	tx := DB.Where("name = ?", name).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

//TODO 用户登录和获取用户信息,密码加密和刷新token
