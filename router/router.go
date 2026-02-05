package router

import (
	"homework_submit/handler"
	"homework_submit/middleware"

	"github.com/jack-wang-176/Maple/web"
)

func Router() *web.HttpService {
	r := web.NewHttpService()
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", handler.UserHandler.Register)
		userGroup.POST("/login", handler.UserHandler.Login)
		userGroup.POST("/refresh", handler.UserHandler.RefreshToken)
	}

	authGroup := r.Group("/")
	authGroup.Use(middleware.AccessTokenDeal)
	{

		authGroup.GET("/submission/excellent", handler.Sub.GetExcellentList) // 查看优秀作业

		student := authGroup.Group("/student")

		student.Use(middleware.CheckStudent)
		{
			// 提交作业
			student.POST("/create", handler.Sub.CreateSub)
			// 查看我的提交
			student.GET("/my", handler.Sub.MySub)
			// 注意：ChangeSub (批改) 不应该在学生组，学生不能改分！
		}

		// --- 管理员接口 (Admin Group) ---
		admin := authGroup.Group("/admin")
		// 嵌入角色检查：只允许 Admin
		admin.Use(middleware.CheckAdmin)
		{
			// 作业管理
			admin.POST("/create", handler.HomeworkHandler.LaunchHomework)
			admin.POST("/delete", handler.HomeworkHandler.DeleteHomework) // 建议改为 DELETE 方法
			admin.POST("/update", handler.HomeworkHandler.UpdateHomework) // 建议改为 PUT 方法

			// 作业查看
			admin.GET("/get", handler.HomeworkHandler.GetHomework)
			admin.GET("/list", handler.HomeworkHandler.GetHomeworkList)

			// 【关键修正】批改作业 (ChangeSub) 应该在这里
			// 只有管理员(老登)才能批改、打分、标记优秀
			admin.POST("/change", handler.Sub.ChangeSub)
		}
	}
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
