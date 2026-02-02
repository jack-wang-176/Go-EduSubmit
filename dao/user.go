package dao

import "homework_submit/model"

type userDao struct{}

var UserDao = new(userDao)

func (d *userDao) CreateUser(u *model.User) error {
	return DB.Create(u).Error
}
func (d *userDao) DeleteUser(u *model.User) error {
	return DB.Delete(u).Error
}

func (d *userDao) GetUserById(id uint) (*model.User, error) {
	var user model.User
	tx := DB.First(&user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (d *userDao) GetUserByName(name string) (*model.User, error) {
	var user model.User
	tx := DB.Where("name = ?", name).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}
func (d *userDao) DetectRole(name string) (bool, error) {
	user, err := UserDao.GetUserByName(name)
	if err != nil {
		return false, err
	}
	if user.Role == model.Admin {
		return true, nil
	}
	return false, nil
}
