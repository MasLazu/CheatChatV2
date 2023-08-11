package service

import (
	"context"
	"errors"

	"github.com/MasLazu/CheatChatV2/model"
	"github.com/MasLazu/CheatChatV2/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(request model.RegisterUserRequest, ctx context.Context) error
}

type UserServiceImpl struct {
}

func NewUserService() UserService {
	return &UserServiceImpl{}
}

func (service UserServiceImpl) Register(request model.RegisterUserRequest, ctx context.Context) error {
	userRepository := repository.NewUsersRepository()
	user, err := userRepository.GetByEmail(ctx, request.Email)
	if err == nil {
		return errors.New("email already used")
	}
	if err != nil && err.Error() != "not found" {
		return errors.New("something went wrong")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("something went wrong")
	}

	user.Email = request.Email
	user.Username = request.Username
	user.Password = string(hashedPassword)

	if err := userRepository.Save(ctx, user); err != nil {
		return errors.New("something went wrong")
	}

	return nil
}
