package service

import (
	"be-assignment/entity"
	"be-assignment/helper"
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

// Send implements TransactionService.
func (t *TransactionServiceImpl) Send(ctx context.Context, fromAccount string, toAccount string, amount float64) error {
	panic("unimplemented")
}

// withdraw implements TransactionService.
func (t *TransactionServiceImpl) withdraw(ctx context.Context, accountNumber string, amount float64) error {
	account, err := t.AccountRepository.FindByAccountNumber(ctx, accountNumber)
	if err != nil {
		return err
	}

	if account.Balance < amount {
		return helper.ErrInsuficientBalance
	}

	trxData := entity.Transaction{
		FromAccount: accountNumber,
		Amount:      amount,
	}

	err = t.TransactionRepository.Save(ctx, trxData)
	if err != nil {
		return err
	}

	accountData := entity.Account{
		Id:      account.Id,
		Balance: account.Balance - amount,
	}

	err = t.AccountRepository.Update(ctx, accountData)
	if err != nil {
		return err
	}

	return nil
}
