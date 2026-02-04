package router

import (
	"homework_submit/handler"
	"homework_submit/middleware"

	"github.com/jack-wang-176/Maple/web"
)

func Router() *web.HttpService {
	r := web.NewHttpService()
	student := r.Group("student")
	admin := r.Group("admin")
	r.POST("/register", handler.UserHandler.Register)
	r.GET("/login", handler.UserHandler.Login)
	//todo 全局中间键

	student.POST("/create", handler.Sub.CreateSub)
	student.GET("/my", handler.Sub.MySub)
	student.POST("/change", handler.Sub.ChangeSub)
	student.GET("/excellent", handler.Sub.GetExcellentList)
	admin.Use(middleware.RoleCheck)
	admin.POST("/create", handler.HomeworkHandler.LaunchHomework)
	admin.POST("/delete", handler.HomeworkHandler.DeleteHomework)
	admin.POST("/update", handler.HomeworkHandler.UpdateHomework)
	admin.GET("/get", handler.HomeworkHandler.GetHomework)
	admin.GET("/list", handler.HomeworkHandler.GetHomeworkList)

	r.Start(":8080")
	return r
}

//func VerifyAccessToken() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		tokenString := c.GetHeader("Authorization")
//		tokenString = model.StripBearer(tokenString)
//		if tokenString == "" {
//			c.JSON(http.StatusUnauthorized, respond.NoToken)
//			c.Abort()
//			return
//		}
//		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//				return nil, respond.WrongTokenMethod
//			}
//			return model.AccessSecret, nil
//		})
//		if err != nil || !token.Valid {
//			c.JSON(http.StatusUnauthorized, respond.InvalidToken)
//			c.Abort()
//			return
//		}
//		if claims, ok := token.Claims.(jwt.MapClaims); ok {
//			tokenType := claims["type"].(string)
//
//			if tokenType == "access" {
//				c.Set("user_id", claims["uid"])
//				c.Set("role", claims["role"])
//				c.Next()
//			} else {
//				c.JSON(http.StatusUnauthorized, respond.InvalidToken)
//				c.Abort()
//				return
//			}
//		}
//
//	}
//}
