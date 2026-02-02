package service

import (
	"errors"
	"homework_submit/dao"
	"homework_submit/model"
	"homework_submit/pkg"

	"gorm.io/gorm"
)

type userService struct{}

var UserService = new(userService)

func (u *userService) Register(username, password, nickname string, role model.Role, dept model.Department) error {
	already, err := dao.UserDao.GetUserByName(username)
	if already == nil || err != nil {
		//TODO 进行错误处理
	}
	harsh, err := pkg.PasswordHarsh(password)
	if err != nil {
		//TODO 进行错误处理
	}
	user := &model.User{
		Name:       username,
		Password:   harsh,
		Nickname:   nickname,
		Role:       role,
		Department: dept,
	}
	err = dao.UserDao.CreateUser(user)
	if err != nil {
		//TODO 进行错误处理
	}
	return nil
}
func (u *userService) Login(username, password string) (string, string, error) {
	user, err := dao.UserDao.GetUserByName(username)
	//TODO 错误处理
	if err != nil {
		// GORM 特有逻辑：如果是 "记录找不到"，这属于业务层面的“用户名错误”
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//TODO 错误处理
		}
		//TODO 错误处理
	}

	if user == nil {
		//TODO 错误处理
	}
	if user.Password == "" {
		//TODO 错误处理
	}
	matched := pkg.DetectPasswordHarsh(password, user.Password)
	if !matched {
		//TODO 错误处理
	}

	return pkg.TokenCreate(user)
}
func (u *userService) DetectUser(username string) (bool, error) {
	role, err := dao.UserDao.DetectRole(username)
	if err != nil {
		return false, err
	}
	return role, nil
}
