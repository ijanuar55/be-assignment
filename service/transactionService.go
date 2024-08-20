package service

import (
	"be-assignment/entity"
	"be-assignment/repository"
	"context"
)

type TransactionServiceImpl struct {
	TransactionRepository repository.TransactionRepository
	AccountRepository     repository.AccountRepository
}

func NewTransactionServiceImpl(transactionRepository repository.TransactionRepository, accountRepository repository.AccountRepository) TransactionService {
	return &TransactionServiceImpl{TransactionRepository: transactionRepository, AccountRepository: accountRepository}
}

// FindByAccountNumber implements TransactionService.
func (t *TransactionServiceImpl) FindByAccountNumber(ctx context.Context, accountNumber string) ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	trx, err := t.TransactionRepository.FindByAccountNumber(ctx, accountNumber)
	if err != nil {
		return transactions, err
	}

	for _, v := range trx {
		transaction := entity.Transaction{
			Id:          v.Id,
			FromAccount: v.FromAccount,
			ToAccount:   v.ToAccount,
			Amount:      v.Amount,
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

// FindById implements TransactionService.
func (t *TransactionServiceImpl) FindById(ctx context.Context, trxId string) (*entity.Transaction, error) {
	trx, err := t.TransactionRepository.FindById(ctx, trxId)
	if err != nil {
		return nil, err
	}
	return &entity.Transaction{
		Id:          trx.Id,
		FromAccount: trx.FromAccount,
		ToAccount:   trx.ToAccount,
		Amount:      trx.Amount,
	}, nil
}

// Save implements TransactionService.
func (t *TransactionServiceImpl) Create(ctx context.Context, req entity.Transaction) error {
	trxData := entity.Transaction{
		Id:          req.Id,
		FromAccount: req.FromAccount,
		ToAccount:   req.ToAccount,
		Amount:      req.Amount,
	}
	err := t.TransactionRepository.Save(ctx, trxData)
	if err != nil {
		return err
	}

	return nil
}
