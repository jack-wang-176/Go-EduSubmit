package handler

import (
	"homework_submit/model"
	"homework_submit/service"
	"net/http"

	"github.com/jack-wang-176/Maple/web"
)

type user struct{}

var UserHandler = &user{}

func (u *user) Login(c *web.Context) error {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var err error
	err = c.BindJson(&req)
	if err != nil {
		//todo
	}
	c.Set("Username", req.Username)
	access, refresh, err := service.UserService.Login(req.Username, req.Password)
	if err != nil {
		//todo
	}
	c.Set("AccessToken", access)
	c.Set("RefreshToken", refresh)
	//todo 完善框架结构
	//todo 返回m结构体
	c.JsonResp()
	return nil
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
		//TODO
	}
	err = service.UserService.Register(req.Username, req.Password, req.NickName, model.Role(req.Role), model.Department(req.Department))
	if err != nil {
		//todo
	}
	//todo
	err = c.JsonResp(http.StatusOK, "注册成功", nil)
	if err != nil {
		//todo
	}

}
