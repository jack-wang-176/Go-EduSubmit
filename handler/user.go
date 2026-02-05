package handler

import (
	"homework_submit/dao"
	"homework_submit/model"
	"homework_submit/pkg"
	"homework_submit/service"
	"time"

	"github.com/jack-wang-176/Maple/web"
)

type user struct{}

var UserHandler = &user{}

func (u *user) Login(c *web.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var err error
	err = c.BindJson(&req)
	if err != nil {
		SendResponse(c, nil, pkg.ParamError)
		return
	}
	c.Set("user", req.Username)

	access, refresh, err := service.UserService.Login(req.Username, req.Password)
	if err != nil {
		SendResponse(c, nil, err)
		return
	}
	user, err := dao.UserDao.GetUserByName(req.Username)
	if err != nil {
		SendResponse(c, nil, err)
		return
	}
	if user == nil {
		SendResponse(c, nil, err)
		return
	}
	expiresAt := time.Now().Add(7 * 24 * time.Hour)
	err = service.UserService.CreateRefresh(user.ID, user.Name, refresh, expiresAt)
	if err != nil {
		SendResponse(c, nil, err)
	}
	SendResponse(c, map[string]interface{}{
		"access_token":  access,
		"refresh_token": refresh,
		"user":          user.ToResponse(),
	}, nil)
}
func (u *user) Register(c *web.Context) {
	var req struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		NickName   string `json:"nickname"`
		Department int8   `json:"department"`
		Role       int8   `json:"role"`
	}
	var err error
	err = c.BindJson(&req)
	if err != nil {
		SendResponse(c, nil, pkg.ParamError)
		return
	}

	err = service.UserService.Register(req.Username, req.Password, req.NickName, model.Role(req.Role), model.Department(req.Department))
	if err != nil {
		SendResponse(c, nil, err)
		return
	}
	SendResponse(c, nil, nil)
}
func (u *user) RefreshToken(c *web.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := c.BindJson(&req); err != nil {
		SendResponse(c, nil, pkg.ParamError)
	}

}
func (u *user) GetProfile(c *web.Context) {
	id, _ := c.Get("userID")
	resp, err := service.UserService.GetProfile(id.(uint))
	if err != nil {
		SendResponse(c, nil, err)
		return
	}
	SendResponse(c, resp, nil)
}
func (u *user) DeleteUser(c *web.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := c.BindJson(&req)
	if err != nil {
		SendResponse(c, nil, pkg.ParamError)
	}

	err = service.UserService.DeleteAccount(req.Username, req.Password)
	if err != nil {
		SendResponse(c, nil, pkg.ErrDeleteUserFailed)
	}
}
