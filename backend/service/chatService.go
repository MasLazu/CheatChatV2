package service

import (
	"context"
	"time"

	"github.com/MasLazu/CheatChatV2/model/domain"
	"github.com/MasLazu/CheatChatV2/repository"
)

type ChatService interface {
	SavePersonalChat(senderEmail string, receiverEmail string, message string, createdAt time.Time) (int64, error)
}

type ChatServiceImpl struct {
}

func NewChatService() ChatService {
	return &ChatServiceImpl{}
}

func (service ChatServiceImpl) SavePersonalChat(senderEmail string, receiverEmail string, message string, createdAt time.Time) (int64, error) {
	ctx := context.TODO()
	var chatId int64

	chatRepository := repository.NewChatsRepository()
	chatRoom, err := chatRepository.GetPersonalChatRoom(ctx, senderEmail, receiverEmail)
	if err != nil {
		return chatId, err
	}

	chat := domain.Chat{
		SenderEmail: senderEmail,
		Message:     message,
		ChatRoom:    chatRoom,
		CreatedAt:   createdAt,
	}

	chatId, err = chatRepository.Save(ctx, chat)
	if err != nil {
		return chatId, err
	}

	return chatId, err
}
