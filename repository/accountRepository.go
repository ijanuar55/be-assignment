package repository

import (
	"be-assignment/entity"
	"be-assignment/helper"
	"be-assignment/prisma/db"
	"context"
	"fmt"
)

type AccountRepositoryImpl struct {
	DB *db.PrismaClient
}

func NewAccountRepository(Db *db.PrismaClient) AccountRepository {
	return &AccountRepositoryImpl{DB: Db}
}

// Delete implements AccountRepository.
func (a *AccountRepositoryImpl) Update(ctx context.Context, account entity.Account) error {
	result, err := a.DB.Account.FindMany(db.Account.ID.Equals(account.Id)).Update(
		db.Account.Balance.Set(account.Balance),
	).Exec(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Rows Affected: ", result)
	return nil
}

// FindById implements AccountRepository.
func (a *AccountRepositoryImpl) FindById(ctx context.Context, accountId string) (*entity.Account, error) {
	account, err := a.DB.Account.FindFirst(db.Account.ID.Equals(accountId)).Exec(ctx)
	if err != nil {
		return nil, err
	}

	accountData := entity.Account{
		Id:            account.ID,
		UserId:        account.UserID,
		Type:          string(account.Type),
		Balance:       account.Balance,
		AccountNumber: account.AccountNumber,
	}

	if account != nil {
		return &accountData, nil
	}

	return &accountData, helper.ErrAccountNotFound
}

func (a *AccountRepositoryImpl) FindByAccountNumber(ctx context.Context, accountNumber string) (*entity.Account, error) {
	account, err := a.DB.Account.FindFirst(db.Account.AccountNumber.Equals(accountNumber)).Exec(ctx)
	if err != nil {
		return nil, err
	}

	accountData := entity.Account{
		Id:            account.ID,
		UserId:        account.UserID,
		Type:          string(account.Type),
		Balance:       account.Balance,
		AccountNumber: account.AccountNumber,
	}

	if account != nil {
		return &accountData, nil
	}

	return nil, helper.ErrAccountNotFound
}

// FindByUserId implements AccountRepository.
func (a *AccountRepositoryImpl) FindByUserId(ctx context.Context, userId string) ([]entity.Account, error) {
	var accounts []entity.Account
	account, err := a.DB.Account.FindMany(db.Account.UserID.Equals(userId)).Exec(ctx)
	if err != nil {
		return accounts, err
	}

	for _, v := range account {
		accountData := entity.Account{
			Id:            v.ID,
			UserId:        v.UserID,
			Type:          string(v.Type),
			Balance:       v.Balance,
			AccountNumber: v.AccountNumber,
		}

		accounts = append(accounts, accountData)
	}

	return accounts, nil
}

// Save implements AccountRepository.
func (a *AccountRepositoryImpl) Save(ctx context.Context, account entity.Account) error {
	result, err := a.DB.Account.CreateOne(
		db.Account.UserID.Set(account.UserId),
		db.Account.Type.Set(db.AccountType(account.Type)),
		db.Account.AccountNumber.Set(account.AccountNumber),
		db.Account.Balance.Set(account.Balance),
	).Exec(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Rows Affected: ", result)
	return nil
}
