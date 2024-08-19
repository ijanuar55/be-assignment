package controller

import (
	"be-assignment/entity"
	"be-assignment/helper"
	"be-assignment/middleware"
	"be-assignment/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (controller *UserController) Create(c *gin.Context) {
	var user entity.UserPostRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": helper.ErrInvalidRequest.Error()})
		return
	}

	signup, err := emailpassword.SignUp(os.Getenv("TENANT_ID"), user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": helper.ErrAlreadyExist.Error()})
		return
	}
	if signup.OK == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": helper.ErrAlreadyExist.Error()})
		return
	}

	dataUser := entity.User{
		Id:    signup.OK.User.ID,
		Email: signup.OK.User.Email,
		Name:  user.Name,
	}

	controller.UserService.Create(c, dataUser)
	c.JSON(http.StatusCreated, entity.WebResponse{
		Message: "User Register Successfully.",
	})
}

func (controller *UserController) Login(c *gin.Context) {
	var loginRequest entity.LoginRequest

	if err := c.BindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": helper.ErrInvalidRequest.Error()})
		return
	}
	signin, err := emailpassword.SignIn(os.Getenv("TENANT_ID"), loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": helper.ErrUnauthorized.Error()})
		return
	}
	if signin.OK == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": helper.ErrUnauthorized.Error()})
		return
	}
	user, err := controller.UserService.FindById(c, signin.OK.User.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": helper.ErrUserNotFound.Error()})
		return
	}

	token, err := middleware.GenerateJWT(user.Email, user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	response := entity.WebResponse{
		Message: "Login Success",
		Data: entity.LoginResponse{
			Id:    user.Id,
			Email: user.Email,
			Name:  user.Name,
			Token: token,
		},
	}
	c.JSON(http.StatusOK, response)
}

func (controller *UserController) GetEmail(c *gin.Context) {
	email, _ := c.Get("email")
	c.JSON(http.StatusOK, gin.H{"message": email})
}
