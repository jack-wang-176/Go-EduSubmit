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
	var err error
	err = c.BindJson(&req)
	if err != nil {
		SendResponse(c, nil, pkg.ParamError)
	}
	c.Set("Username", req.Username)
	access, refresh, err := service.UserService.Login(req.Username, req.Password)
	if err != nil {
		SendResponse(c, nil, pkg.ParamError)
	}
	c.Set("AccessToken", access)
	c.Set("RefreshToken", refresh)
	SendResponse(c, map[string]string{
		"access_token":  access,
		"refresh_token": refresh,
	}, pkg.Success)
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
	}
	err = service.UserService.Register(req.Username, req.Password, req.NickName, model.Role(req.Role), model.Department(req.Department))
	if err != nil {
		SendResponse(c, nil, pkg.ServerError)
	}
	SendResponse(c, nil, pkg.Success)

}
