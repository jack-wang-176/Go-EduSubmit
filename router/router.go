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
		userGroup.POST("/refresh", handler.Token.RefreshToken)
	}

	authGroup := r.Group("/")
	authGroup.Use(middleware.AccessTokenDeal) // 全局 JWT 校验
	{

		authGroup.GET("/user/profile", handler.UserHandler.GetProfile)
		authGroup.DELETE("/user/account", handler.UserHandler.DeleteUser)

		// 作业模块 (Homework)
		homework := authGroup.Group("/homework")
		{
			homework.GET("", handler.HomeworkHandler.GetHomeworkList)
			homework.GET("/:id", handler.HomeworkHandler.GetHomework)

			// 管理员专属
			homework.POST("", middleware.CheckAdmin, handler.HomeworkHandler.LaunchHomework)
			homework.PUT("/:id", middleware.CheckAdmin, handler.HomeworkHandler.UpdateHomework)
			homework.DELETE("/:id", middleware.CheckAdmin, handler.HomeworkHandler.DeleteHomework)
		}

		submission := authGroup.Group("/submission")
		{
			submission.GET("/excellent", handler.Sub.GetExcellentList)

			submission.POST("", middleware.CheckStudent, handler.Sub.CreateSub)
			submission.GET("/my", middleware.CheckStudent, handler.Sub.MySub)

			submission.GET("/homework/:id", middleware.CheckAdmin, handler.Sub.GetWorkSubs)
			submission.PUT("/:id/review", middleware.CheckAdmin, handler.Sub.ChangeSub)
			submission.PUT("/:id/excellent", middleware.CheckAdmin, handler.Sub.MarkExcellent)
		}
	}
	return r
}
