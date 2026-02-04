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
	if already != nil || err != nil {
		return pkg.ErrUserExists
	}
	harsh, err := pkg.PasswordHarsh(password)
	if err != nil {
		return pkg.ErrPasswordIncorrect
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
		return pkg.ErrorPkg.WithCause(err)
	}
	return nil
}
func (u *userService) Login(username, password string) (string, string, error) {
	user, err := dao.UserDao.GetUserByName(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", "", pkg.ErrUserNotFound
		}
		return "", "", pkg.ErrorPkg.WithCause(err)
	}

	if user == nil {
		return "", "", pkg.NoInput
	}
	if user.Password == "" {
		return "", "", pkg.NoInput
	}
	matched := pkg.DetectPasswordHarsh(password, user.Password)
	if !matched {
		return "", "", pkg.ErrPasswordIncorrect
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
