package routes

import (
	"be-assignment/handlers"
	"be-assignment/middleware"

	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
)

func Router() *gin.Engine {
	r := gin.Default()

	// // Define routes and use handler functions
	// r.GET("/ping", handlers.PingHandler)

	api := r.Group("/api")
	{
		// Protected route example
		sessionRequired := false
		api.GET("/user/:email", middleware.VerifySession(&sessmodels.VerifySessionOptions{
			SessionRequired: &sessionRequired,
		}), handlers.GetEmail)
	}

	// Add SuperTokens API routes
	r.POST("/auth/signup", handlers.EmailPasswordSignUp)
	r.POST("/auth/login", handlers.EmailPasswordSignIn)
	// r.POST("/auth/refresh", session.RefreshSession)
	// r.POST("/auth/logout", session.RevokeSession)

	return r
}
