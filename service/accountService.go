package service

import (
	"be-assignment/entity"
	"be-assignment/repository"
	"context"
)

type AccountServiceImpl struct {
	AccountRepository repository.AccountRepository
}

func NewAccountServiceImpl(accountRepository repository.AccountRepository) AccountService {
	return &AccountServiceImpl{AccountRepository: accountRepository}
}

// Create implements AccountService.
func (a *AccountServiceImpl) Create(ctx context.Context, req entity.Account) error {
	accountData := entity.Account{
		Id:            req.Id,
		UserId:        req.UserId,
		Type:          req.Type,
		Balance:       req.Balance,
		AccountNumber: req.AccountNumber,
	}
	a.AccountRepository.Save(ctx, accountData)
	return nil
}

// FindById implements AccountService.
func (a *AccountServiceImpl) FindById(ctx context.Context, accountId string) (*entity.Account, error) {
	account, err := a.AccountRepository.FindById(ctx, accountId)
	if err != nil {
		return nil, err
	}
	return &entity.Account{
		Id:            account.Id,
		UserId:        account.UserId,
		Type:          account.Type,
		Balance:       account.Balance,
		AccountNumber: account.AccountNumber,
	}, nil
}

// FindByUserId implements AccountService.
func (a *AccountServiceImpl) FindByUserId(ctx context.Context, userId string) ([]entity.Account, error) {
	var accounts []entity.Account
	acc, err := a.AccountRepository.FindByUserId(ctx, userId)
	if err != nil {
		return accounts, err
	}

	for _, v := range acc {
		account := entity.Account{
			Id:            v.Id,
			UserId:        v.UserId,
			Type:          v.Type,
			Balance:       v.Balance,
			AccountNumber: v.AccountNumber,
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}
