package handler

import (
	"homework_submit/pkg"
	"homework_submit/service"

	"github.com/jack-wang-176/Maple/web"
)

type token struct{}

var Token token

func (t *token) RefreshToken(c *web.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.BindJson(&req); err != nil {
		SendResponse(c, nil, pkg.ParamError)
		return
	}

	newAccess, newRefresh, err := service.UserService.RefreshToken(req.RefreshToken)
	if err != nil {
		SendResponse(c, nil, err)
		return
	}

	SendResponse(c, map[string]string{
		"access_token":  newAccess,
		"refresh_token": newRefresh,
	}, nil, "刷新成功")
}
