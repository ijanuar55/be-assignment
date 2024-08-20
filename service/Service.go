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
	Update(ctx context.Context, req entity.Account) error
	FindByUserId(ctx context.Context, userId string) ([]entity.Account, error)
	FindById(ctx context.Context, accountId string) (*entity.Account, error)
	FindByAccountNumber(ctx context.Context, accountNumber string) (*entity.Account, error)
}

type TransactionService interface {
	Create(ctx context.Context, req entity.Transaction) error
	FindById(ctx context.Context, trxId string) (*entity.Transaction, error)
	FindByAccountNumber(ctx context.Context, accountNumber string) ([]entity.Transaction, error)
}
