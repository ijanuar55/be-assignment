package routes

import (
	"be-assignment/controller"
	"be-assignment/middleware"

	"github.com/gin-gonic/gin"
)

func Router(user *controller.UserController, account *controller.AccountController) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	api.Use(middleware.AuthenticateJWT())
	{
		api.GET("/user/:email", user.GetEmail)
		api.GET("/account/:userId", account.FindByUserId)
		api.POST("/account", account.Create)
	}

	r.POST("/auth/signup", user.Create)
	r.POST("/auth/login", user.Login)

	return r
}
