package service

import (
	"context"
	"errors"
	"github.com/MasLazu/CheatChatV2/model/web"

	"github.com/MasLazu/CheatChatV2/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(request web.RegisterUserRequest, ctx context.Context) error
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (service UserServiceImpl) Register(request web.RegisterUserRequest, ctx context.Context) error {
	user, err := service.userRepository.GetByEmail(ctx, request.Email)
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

	if err := service.userRepository.Save(ctx, user); err != nil {
		return errors.New("something went wrong")
	}

	return nil
}
