package service

import (
	"homework_submit/dao"
	"homework_submit/pkg"
	"time"
)

func (u *userService) RefreshToken(oldRefreshToken string) (newAccess string, newRefresh string, err error) {

	claims, err := pkg.ParseRefreshToken(oldRefreshToken)
	if err != nil {
		return "", "", pkg.ErrorPkg.WithCause(err)
	}
	token, err := dao.Refresh.GetValidToken(oldRefreshToken)
	if err != nil || token == nil {
		return "", "", pkg.TokenFailed
	}
	user, err := dao.UserDao.GetUserById(claims.UserID)
	if err != nil {
		return "", "", pkg.ErrUserNotFound
	}
	err = dao.Refresh.Revoke(oldRefreshToken)
	if err != nil {
		return "", "", err
	}
	newAccess, newRefresh, err = pkg.TokenCreate(user)
	if err != nil {
		return "", "", pkg.ErrorPkg.WithCause(err)
	}
	newExpiresAt := time.Now().Add(7 * 24 * time.Hour)
	err = dao.Refresh.Create(user.ID, user.Name, newRefresh, newExpiresAt)
	if err != nil {
		return "", "", pkg.ErrorPkg.WithCause(err)
	}
	return newAccess, newRefresh, nil
}
