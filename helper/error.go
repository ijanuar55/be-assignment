package helper

import "errors"

var (
	ErrInvalidRequest      = errors.New("invalid request")
	ErrAlreadyExist        = errors.New("email already exist")
	ErrUnauthorized        = errors.New("wrong email / password")
	ErrUserNotFound        = errors.New("user not found")
	ErrAccountNotFound     = errors.New("account not found")
	ErrTransactionNotFound = errors.New("transaction not found")
	ErrInsuficientBalance  = errors.New("insuficient balance")
	ErrDontHaveAccess      = errors.New("you dont have access")
	ErrCannotCreateTran    = errors.New("cannot create transaction")
	ErrCannotUpdateAcc     = errors.New("cannot update account")
)

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}
