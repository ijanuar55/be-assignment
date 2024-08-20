package routes

import (
	"be-assignment/controller"
	"be-assignment/middleware"

	"github.com/gin-gonic/gin"
)

func Router(user *controller.UserController, account *controller.AccountController, transaction *controller.TransactionController) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	api.Use(middleware.AuthenticateJWT())
	{
		api.GET("/user/:email", user.GetEmail)
		api.GET("/account/:userId", account.FindByUserId)
		api.GET("/transaction/:accNumber", transaction.GetTransactionByAccountNumber)
		api.POST("/account", account.Create)
		api.POST("/send", transaction.Send)
		api.POST("/withdraw", transaction.Withdraw)
	}

	r.POST("/auth/signup", user.Create)
	r.POST("/auth/login", user.Login)

	return r
}
