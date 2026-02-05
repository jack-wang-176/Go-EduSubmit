package middleware

import (
	"homework_submit/handler"
	"homework_submit/pkg"
	"strings"

	"github.com/jack-wang-176/Maple/web"
)

func AccessTokenDeal(c *web.Context) {
	//todo 换成context内部封装的方法
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		handler.SendResponse(c, nil, pkg.TokenErr)
		c.Abort()
		return
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		handler.SendResponse(c, nil, pkg.TokenErr)
		c.Abort()
		return
	}
	claim, err := pkg.ParseAccessToken(parts[1])
	if err != nil {
		handler.SendResponse(c, nil, pkg.TokenFailed)
		c.Abort()
		return
	}
	c.Set("userID", claim.UserID)
	c.Set("user", claim.Name)
	c.Set("role", claim.Role)
	c.Set("department", claim.Department)
	c.Next()
}
