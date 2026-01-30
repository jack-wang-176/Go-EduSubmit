package user

import (
	"homework_submit/dao"
	"homework_submit/model"
)

func TheUserCreate(user *model.User) error {
	tx := dao.DB.Create(user)
	if tx.Error != nil || tx.RowsAffected <= 0 {
		return tx.Error
	}
	return nil
}
func TheUserDelete(user *model.User) error {
	tx := dao.DB.Delete(user)
	if tx.Error != nil || tx.RowsAffected <= 0 {
		return tx.Error
	}
	return nil
}
