package service

import (
	"be-assignment/entity"
	"be-assignment/repository"
	"context"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

// Create implements UserService.
func (u *UserServiceImpl) Create(ctx context.Context, request entity.User) error {
	userData := entity.User{
		Id:    request.Id,
		Email: request.Email,
		Name:  request.Name,
	}
	u.UserRepository.Save(ctx, userData)
	return nil
}

// Delete implements UserService.
func (u *UserServiceImpl) Delete(ctx context.Context, userId string) error {
	user, err := u.UserRepository.FindById(ctx, userId)
	if err != nil {
		return err
	}
	u.UserRepository.Delete(ctx, user.Id)
	return nil
}

// FindById implements UserService.
func (u *UserServiceImpl) FindById(ctx context.Context, userId string) (*entity.UserResponse, error) {
	user, err := u.UserRepository.FindById(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &entity.UserResponse{
		Id:    user.Id,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}

// Update implements UserService.
func (u *UserServiceImpl) Update(ctx context.Context, request entity.UserUpdateRequest) error {
	u.UserRepository.Update(ctx, entity.User{
		Id:    request.Id,
		Email: request.Email,
		Name:  request.Name,
	})
	return nil
}
