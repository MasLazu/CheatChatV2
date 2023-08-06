package websocketProvider

import (
	"github.com/MasLazu/CheatChatV2/model"
	"github.com/MasLazu/CheatChatV2/service"
	"log"
	"time"
)

func (manager *Manager) SendPersonalChatController(messageRequest map[string]any) {
	senderEmailReq, ok := messageRequest["sender_email"].(string)
	if !ok {
		log.Println("sender_email error")
		return
	}

	messageReq, ok := messageRequest["message"].(string)
	if !ok {
		log.Println("message error")
		return
	}

	receiverEmailReq, ok := messageRequest["receiver_email"].(string)
	if !ok {
		log.Println("receiver_email")
		return
	}

	message := model.ChatResponse{
		SenderEmail: senderEmailReq,
		Message:     messageReq,
		CreatedAt:   time.Now(),
	}

	chatService := service.NewChatService()
	id, err := chatService.SavePersonalChat(senderEmailReq, receiverEmailReq, messageReq, message.CreatedAt)
	if err != nil {
		return
	}

	message.Id = id
	manager.SendMessageToUser(receiverEmailReq, message)
}
