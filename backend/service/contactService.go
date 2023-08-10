package service

import (
	"context"
	"errors"
	"github.com/MasLazu/CheatChatV2/model/domain"
	"github.com/MasLazu/CheatChatV2/repository"
)

type ContactService interface {
	AddContact(request domain.Contact, ctx context.Context) error
}

type ContactServiceImpl struct {
}

func NewContactService() ContactService {
	return &ContactServiceImpl{}
}

func (service ContactServiceImpl) AddContact(request domain.Contact, ctx context.Context) error {
	userRepository := repository.NewUsersRepository()
	if _, err := userRepository.GetByEmail(ctx, request.SavedUserEmail); err != nil {
		return errors.New("user not found")
	}

	personalRepository := repository.NewPersonalRepository()
	if _, err := personalRepository.GetByMember(request.UserEmail, request.SavedUserEmail, ctx); err != nil {
		if err.Error() == "not found" {
			if err := personalRepository.Save(request.UserEmail, request.SavedUserEmail, ctx); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	contactRepository := repository.NewContactRepository()
	if err := contactRepository.Save(ctx, request); err != nil {
		return err
	}

	return nil
}
