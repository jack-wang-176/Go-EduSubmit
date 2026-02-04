package middleware

import (
	"homework_submit/handler"
	"homework_submit/pkg"
	"homework_submit/service"

	"github.com/jack-wang-176/Maple/web"
)

func RoleCheck(c *web.Context) {
	user, flag := c.Get("user")
	if !flag || user == nil {
		handler.SendResponse(c, nil, pkg.ServerError)
		return
	}
	detectUser, err := service.UserService.DetectUser(user.(string))
	if err != nil {
		handler.SendResponse(c, nil, err)
		return
	}

	if !detectUser {
		handler.SendResponse(c, nil, pkg.ErrUserNotAdmin)
		c.Abort()
		return
	}
	c.Next()
}
