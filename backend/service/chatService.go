package service

import (
	"context"
	"time"

	"github.com/MasLazu/CheatChatV2/model/domain"
	"github.com/MasLazu/CheatChatV2/repository"
)

type ChatService struct {
	chatRepository     *repository.ChatRepository
	personalRepository *repository.PersonalRepository
	groupRepository    *repository.GroupRepository
}

func NewChatService(chatRepository *repository.ChatRepository, personalRepository *repository.PersonalRepository, groupRepository *repository.GroupRepository) *ChatService {
	return &ChatService{
		chatRepository:     chatRepository,
		personalRepository: personalRepository,
		groupRepository:    groupRepository,
	}
}

func (service *ChatService) SavePersonalChat(senderEmail string, receiverEmail string, message string, createdAt time.Time) (int64, error) {
	ctx := context.TODO()
	var chatId int64

	chatRoom, err := service.personalRepository.GetChatRoom(ctx, senderEmail, receiverEmail)
	if err != nil {
		return chatId, err
	}

	chat := domain.Chat{
		SenderEmail: senderEmail,
		Message:     message,
		ChatRoom:    chatRoom,
		CreatedAt:   createdAt,
	}

	chatId, err = service.chatRepository.Save(ctx, chat)
	if err != nil {
		return chatId, err
	}

	return chatId, err
}

func (service *ChatService) SaveGroupChat(senderEmail string, groupId int64, message string, createdAt time.Time) (int64, error) {
	ctx := context.TODO()
	var chatId int64

	chatRoom, err := service.groupRepository.GetChatRoom(ctx, groupId)
	if err != nil {
		return chatId, err
	}

	chat := domain.Chat{
		SenderEmail: senderEmail,
		Message:     message,
		ChatRoom:    chatRoom,
		CreatedAt:   createdAt,
	}

	chatId, err = service.chatRepository.Save(ctx, chat)
	if err != nil {
		return chatId, err
	}

	return chatId, err
}
