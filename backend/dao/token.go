package dao

import (
	"errors"
	"homework_submit/model"
	"time"

	"gorm.io/gorm"
)

type tokens struct{}

var Refresh = new(tokens)

func (d *tokens) Create(userID uint, user, token string, expiresAt time.Time) error {
	rt := model.RefreshToken{
		UserID:    userID,
		UserName:  user,
		Token:     token,
		ExpiresAt: expiresAt,
		Revoked:   false,
	}
	return DB.Create(&rt).Error
}

func (d *tokens) GetValidToken(tokenStr string) (*model.RefreshToken, error) {
	var rt model.RefreshToken

	err := DB.Where("token = ? AND revoked = ?", tokenStr, false).First(&rt).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("refresh token 不存在或已失效")
		}
		return nil, err
	}

	if rt.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("refresh token 已过期")
	}

	return &rt, nil
}

func (d *tokens) Revoke(tokenStr string) error {
	return DB.Model(&model.RefreshToken{}).
		Where("token = ?", tokenStr).
		Update("revoked", true).Error
}

func (d *tokens) DeleteByUserID(userID uint) error {
	return DB.Where("user_id = ?", userID).Delete(&model.RefreshToken{}).Error
}
