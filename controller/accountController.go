package controller

import (
	"be-assignment/entity"
	"be-assignment/helper"
	"be-assignment/service"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	AccountService service.AccountService
	UserService    service.UserService
}

func NewAccountController(accountService service.AccountService, userService service.UserService) *AccountController {
	return &AccountController{AccountService: accountService, UserService: userService}
}

func (controller *AccountController) Create(c *gin.Context) {
	var account entity.AccountRequest

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": helper.ErrInvalidRequest.Error()})
		return
	}

	_, err := controller.UserService.FindById(c, account.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": helper.ErrUserNotFound.Error()})
		return
	}

	var accountNumber string

	switch account.Type {
	case "debit":
		accountNumber += "001"
	case "credit":
		accountNumber += "002"
	case "loan":
		accountNumber += "003"
	default:
		accountNumber += "000"
	}

	year := strconv.Itoa(time.Now().Year())
	month := int(time.Now().Month())
	date := int(time.Now().Day())
	hour := strconv.Itoa(time.Now().Hour())
	minute := strconv.Itoa(time.Now().Minute())
	second := strconv.Itoa(time.Now().Second())
	accountNumber += year + fmt.Sprintf("%02d", month) + fmt.Sprintf("%02d", date) + hour + minute + second

	dataAccount := entity.Account{
		UserId:        account.UserId,
		Type:          account.Type,
		Balance:       account.Balance,
		AccountNumber: accountNumber,
	}
	controller.AccountService.Create(c, dataAccount)

	c.JSON(http.StatusCreated, entity.WebResponse{
		Message: "Account Created Successfully.",
	})
}

func (controller *AccountController) FindByUserId(c *gin.Context) {
	userId := c.Param("userId")

	result, _ := controller.AccountService.FindByUserId(c, userId)

	c.JSON(http.StatusOK, entity.WebResponse{
		Message: "Account Retrieve Successfully.",
		Data:    result,
	})
}
