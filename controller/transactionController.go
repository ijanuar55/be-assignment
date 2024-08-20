package controller

import (
	"be-assignment/entity"
	"be-assignment/helper"
	"be-assignment/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TransactionService service.TransactionService
	AccountService     service.AccountService
}

func NewTransactionController(transactionService service.TransactionService, accountService service.AccountService) *TransactionController {
	return &TransactionController{TransactionService: transactionService, AccountService: accountService}
}

func (controller *TransactionController) GetTransactionByAccountNumber(c *gin.Context) {
	accountNumber := c.Param("accNumber")
	userId := c.Value("user_id")
	fmt.Println(userId)

	account, err := controller.AccountService.FindByAccountNumber(c, accountNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": helper.ErrAccountNotFound.Error()})
		return
	}
	if account.UserId != userId {
		c.JSON(http.StatusForbidden, gin.H{"message": helper.ErrDontHaveAccess.Error()})
		return
	}

	trx, _ := controller.TransactionService.FindByAccountNumber(c, accountNumber)

	c.JSON(http.StatusOK, entity.WebResponse{Message: "Success", Data: trx})
}

func (controller *TransactionController) Send(c *gin.Context) {
	userId := c.Value("user_id")
	var transaction entity.Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": helper.ErrInvalidRequest.Error()})
		return
	}

	trxData := entity.Transaction{
		FromAccount: transaction.FromAccount,
		ToAccount:   transaction.ToAccount,
		Amount:      transaction.Amount,
	}

	accounts := []string{transaction.FromAccount, transaction.ToAccount}
	for i, v := range accounts {
		var balanceNow float64
		account, err := controller.AccountService.FindByAccountNumber(c, v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": helper.ErrAccountNotFound.Error()})
			return
		}
		if i == 0 {
			if account.UserId != userId {
				c.JSON(http.StatusForbidden, gin.H{"message": helper.ErrDontHaveAccess.Error()})
				return
			}
			if account.Balance < transaction.Amount {
				c.JSON(http.StatusBadRequest, gin.H{"message": helper.ErrInsuficientBalance.Error()})
				return
			}
			balanceNow = account.Balance - transaction.Amount
			fmt.Println(balanceNow)
		}
		balanceNow = account.Balance + transaction.Amount
		accData := entity.Account{
			Id:      account.Id,
			Balance: balanceNow,
		}
		err = controller.AccountService.Update(c, accData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": helper.ErrCannotUpdateAcc.Error()})
			return
		}
	}

	err := controller.TransactionService.Create(c, trxData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": helper.ErrCannotCreateTran.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

// withdraw implements TransactionService.
func (controller *TransactionController) Withdraw(c *gin.Context) {
	var transaction entity.Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": helper.ErrInvalidRequest.Error()})
		return
	}
	account, err := controller.AccountService.FindByAccountNumber(c, transaction.FromAccount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": helper.ErrAccountNotFound.Error()})
		return
	}

	if account.Balance < transaction.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"message": helper.ErrInsuficientBalance.Error()})
		return
	}

	trxData := entity.Transaction{
		FromAccount: transaction.FromAccount,
		Amount:      transaction.Amount,
	}

	err = controller.TransactionService.Create(c, trxData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": helper.ErrCannotCreateTran.Error()})
		return
	}

	accountData := entity.Account{
		Id:      account.Id,
		Balance: account.Balance - transaction.Amount,
	}

	err = controller.AccountService.Update(c, accountData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": helper.ErrCannotUpdateAcc.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}
