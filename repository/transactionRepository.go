package repository

import (
	"be-assignment/entity"
	"be-assignment/helper"
	"be-assignment/prisma/db"
	"context"
	"fmt"
)

type TransactionRepositoryImpl struct {
	DB *db.PrismaClient
}

func NewTransactionRepository(Db *db.PrismaClient) TransactionRepository {
	return &TransactionRepositoryImpl{DB: Db}
}

func (t *TransactionRepositoryImpl) FindById(ctx context.Context, trxId string) (*entity.Transaction, error) {
	trx, err := t.DB.Transaction.FindFirst(db.Transaction.ID.Equals(trxId)).Exec(ctx)
	if err != nil {
		return nil, err
	}

	trxData := entity.Transaction{
		Id:          trx.ID,
		FromAccount: trx.FromAccount,
		ToAccount:   trx.ToAccount,
		Amount:      trx.Amount,
	}

	if trx != nil {
		return &trxData, nil
	}

	return nil, helper.ErrTransactionNotFound
}

// FindByAccountNumber implements TransactionRepository.
func (t *TransactionRepositoryImpl) FindByAccountNumber(ctx context.Context, accountNumber string) ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	trx, err := t.DB.Transaction.FindMany(db.Transaction.FromAccount.Equals(accountNumber)).Exec(ctx)
	if err != nil {
		return transactions, err
	}

	for _, v := range trx {
		trxData := entity.Transaction{
			Id:          v.ID,
			FromAccount: v.FromAccount,
			ToAccount:   v.ToAccount,
			Amount:      v.Amount,
		}

		transactions = append(transactions, trxData)
	}

	return transactions, nil
}

func (t *TransactionRepositoryImpl) Save(ctx context.Context, trx entity.Transaction) error {
	result, err := t.DB.Transaction.CreateOne(
		db.Transaction.FromAccount.Set(trx.FromAccount),
		db.Transaction.ToAccount.Set(trx.ToAccount),
		db.Transaction.Amount.Set(trx.Amount),
	).Exec(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Rows Affected: ", result)
	return nil
}
