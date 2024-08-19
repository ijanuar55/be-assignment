package repository

import (
	"be-assignment/entity"
	"context"
)

type UserRepository interface {
	Save(ctx context.Context, user entity.User) error
	Update(ctx context.Context, user entity.User) error
	Delete(ctx context.Context, userId string) error
	FindById(ctx context.Context, userId string) (*entity.User, error)
}

type AccountRepository interface {
	Save(ctx context.Context, account entity.Account) error
	Update(ctx context.Context, account entity.Account) error
	FindByUserId(ctx context.Context, userId string) ([]entity.Account, error)
	FindById(ctx context.Context, accountId string) (*entity.Account, error)
	FindByAccountNumber(ctx context.Context, accountNumber string) (*entity.Account, error)
}

type TransactionRepository interface {
	Save(ctx context.Context, trx entity.Transaction) error
	FindById(ctx context.Context, trxId string) (*entity.Transaction, error)
	FindByAccountNumber(ctx context.Context, accountNumber string) ([]entity.Transaction, error)
}
