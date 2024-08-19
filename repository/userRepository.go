package repository

import (
	"be-assignment/entity"
	"be-assignment/helper"
	"be-assignment/prisma/db"
	"context"
	"fmt"
)

type UserRepositoryImpl struct {
	DB *db.PrismaClient
}

func NewUserRepository(Db *db.PrismaClient) UserRepository {
	return &UserRepositoryImpl{DB: Db}
}

func (p *UserRepositoryImpl) Delete(ctx context.Context, userId string) error {
	result, err := p.DB.User.FindUnique(db.User.ID.Equals(userId)).Delete().Exec(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Rows Affected: ", result)
	return nil
}

func (p *UserRepositoryImpl) Save(ctx context.Context, user entity.User) error {
	result, err := p.DB.User.CreateOne(
		db.User.ID.Set(user.Id),
		db.User.Email.Set(user.Email),
		db.User.Name.Set(user.Name),
	).Exec(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Rows Affected: ", result)
	return nil
}

func (p *UserRepositoryImpl) Update(ctx context.Context, user entity.User) error {
	result, err := p.DB.User.FindMany(db.User.ID.Equals(user.Id)).Update(
		db.User.Email.Set(user.Email),
		db.User.Name.Set(user.Name),
	).Exec(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Rows Affected: ", result)
	return nil
}

func (p *UserRepositoryImpl) FindById(ctx context.Context, userId string) (*entity.User, error) {
	user, err := p.DB.User.FindFirst(db.User.ID.Equals(userId)).Exec(ctx)
	if err != nil {
		return nil, helper.ErrUserNotFound
	}

	userData := entity.User{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	if user != nil {
		return &userData, nil
	}

	return &userData, helper.ErrUserNotFound
}
