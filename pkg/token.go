package pkg

import (
	"fmt"
	"homework_submit/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	AccessSecret  = []byte("access_secret_example_change_me")
	RefreshSecret = []byte("refresh_secret_example_change_me")
)

//type User struct {
//	gorm.Model
//	Name       string     `gorm:"type:varchar(50);not null;unique" json:"name"`
//	Email      string     `gorm:"type:varchar(100);not null;unique" json:"email"`
//	Password   string     `gorm:"type:varchar(255);not null" json:"-"`
//	Nickname   string     `gorm:"type:varchar(50);not null" json:"nickname"`
//	Role       Role       `gorm:"type:tinyint;not null;default:1;comment:1=Student,2=Admin" json:"role"`
//	Department Department `gorm:"type:tinyint;not null;default:1;comment:1=Backend..." json:"department"`
//}

type tokenClaim struct {
	UserID     uint             `gorm:"not null;unique" json:"user_id"`
	Name       string           `gorm:"type:varchar(50);not null;unique" json:"name"`
	Email      string           `gorm:"type:varchar(100);not null;unique" json:"email"`
	Role       model.Role       `gorm:"type:tinyint;not null;default:1;comment:1=Student,2=Admin" json:"role"`
	Department model.Department `gorm:"type:tinyint;not null;default:1;comment:1=Backend..." json:"department"`
	jwt.RegisteredClaims
}

func TokenCreate(user *model.User) (accessToken string, refreshToken string, err error) {
	now := time.Now()
	accessClaim := tokenClaim{
		UserID:     user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Role:       user.Role,
		Department: user.Department,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 2)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaim)
	accessToken, err = claims.SignedString(AccessSecret)
	if err != nil {
		return "", "", err
	}
	refreshClaim := tokenClaim{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Minute * 10)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	refreshTok := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaim)
	refreshToken, err = refreshTok.SignedString(RefreshSecret)
	if err != nil {
		return "", "", fmt.Errorf("在验证refreshtoken时遭到失败:%w", err)
	}
	return accessToken, refreshToken, nil

}
