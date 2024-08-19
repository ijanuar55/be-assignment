package service

import (
	"be-assignment/entity"
	"context"
)

type UserService interface {
	Create(ctx context.Context, request entity.User) error
	Update(ctx context.Context, request entity.UserUpdateRequest) error
	Delete(ctx context.Context, userId string) error
	FindById(ctx context.Context, userId string) (*entity.UserResponse, error)
}

type AccountService interface {
	Create(ctx context.Context, req entity.Account) error
	FindByUserId(ctx context.Context, userId string) ([]entity.Account, error)
	FindById(ctx context.Context, accountId string) (*entity.Account, error)
}

type TransactionService interface {
	Send(ctx context.Context, fromAccount, toAccount string, amount float64) error
	withdraw(ctx context.Context, accountNumber string, amount float64) error
}
