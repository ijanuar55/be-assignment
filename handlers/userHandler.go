package handlers

import (
	"be-assignment/entity"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
)

const tenantId string = "public"

func EmailPasswordSignUp(c *gin.Context) {
	var loginRequest entity.LoginRequest

	if err := c.BindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	response, err := emailpassword.SignUp(tenantId, loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func EmailPasswordSignIn(c *gin.Context) {
	var loginRequest entity.LoginRequest

	if err := c.BindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	response, err := emailpassword.SignIn(tenantId, loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func GetEmail(c *gin.Context) {
	sessionContainer := session.GetSessionFromRequestContext(c.Request.Context())
	userID := sessionContainer.GetUserID()
	fmt.Println(userID)
	email := c.Param("email")

	userInfo, err := emailpassword.GetUserByEmail(os.Getenv("TENANT_ID"), email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, userInfo)
}
