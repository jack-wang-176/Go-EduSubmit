package handler

import (
	"homework_submit/model"
	"homework_submit/pkg"
	"homework_submit/service"

	"github.com/jack-wang-176/Maple/web"
)

type user struct{}

var UserHandler = &user{}

func (u *user) Login(c *web.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJson(&req); err != nil {
		SendResponse(c, nil, pkg.ParamError)
		return
	}

	access, refresh, err := service.UserService.Login(req.Username, req.Password)
	if err != nil {
		SendResponse(c, nil, err)
		return
	}

	user, err := service.UserService.GetAccount(req.Username)
	if err != nil {
		SendResponse(c, nil, err)
		return
	}

	SendResponse(c, map[string]interface{}{
		"access_token":  access,
		"refresh_token": refresh,
		"user":          user.ToResponse(),
	}, nil, "登录成功")
}
func (u *user) Register(c *web.Context) {
	var req struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		NickName   string `json:"nickname"`
		Department string `json:"department"`
		Role       string `json:"role"`
	}
	var err error
	err = c.BindJson(&req)
	if err != nil {
		SendResponse(c, nil, pkg.ParamError)
		return
	}

	department := model.Depart[req.Department]
	role := model.Roles[req.Role]
	theUser, err := service.UserService.Register(req.Username, req.Password, req.NickName, role, department)
	if err != nil {
		SendResponse(c, nil, err)
		return
	}
	SendResponse(c, theUser.ToResponse(), nil, "注册成功")
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
	user, err := service.UserService.GetProfile(id.(uint))
	if err != nil {
		SendResponse(c, nil, err)
		return
	}
	SendResponse(c, user.ToResponse(), nil)
}
func (u *user) DeleteUser(c *web.Context) {
	var req struct {
		Password string `json:"password"`
	}
	id, _ := c.Get("userID")
	err := c.BindJson(&req)
	if err != nil {
		SendResponse(c, nil, pkg.ParamError)
		return
	}

	err = service.UserService.DeleteAccount(id.(uint), req.Password)
	if err != nil {
		SendResponse(c, nil, pkg.ErrDeleteUserFailed)
		return
	}
	SendResponse(c, nil, nil, "账号已注销")
}
