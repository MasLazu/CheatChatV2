package service

import (
	"context"
	"errors"

	"github.com/MasLazu/CheatChatV2/model/domain"
	"github.com/MasLazu/CheatChatV2/repository"
)

type ContactService struct {
	userRepository     *repository.UserRepository
	personalRepository *repository.PersonalRepository
	contactRepository  *repository.ContactRepository
}

func NewContactService(userRepository *repository.UserRepository, personalRepository *repository.PersonalRepository, contactRepository *repository.ContactRepository) *ContactService {
	return &ContactService{
		userRepository:     userRepository,
		personalRepository: personalRepository,
		contactRepository:  contactRepository,
	}
}

func (service *ContactService) AddContact(request domain.Contact, ctx context.Context) error {
	if _, err := service.userRepository.GetByEmail(ctx, request.SavedUserEmail); err != nil {
		return errors.New("user not found")
	}

	if _, err := service.personalRepository.GetByMember(request.UserEmail, request.SavedUserEmail, ctx); err != nil {
		if err.Error() == "not found" {
			if err := service.personalRepository.Save(request.UserEmail, request.SavedUserEmail, ctx); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	if err := service.contactRepository.Save(ctx, request); err != nil {
		return err
	}

	return nil
}
