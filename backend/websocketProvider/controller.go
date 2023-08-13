package websocketProvider

import (
	"log"
	"time"

	"github.com/MasLazu/CheatChatV2/model/web"
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

	message := web.ChatResponse{
		SenderEmail:   senderEmailReq,
		ReceiverEmail: receiverEmailReq,
		Message:       messageReq,
		CreatedAt:     time.Now(),
	}

	log.Println("SendPersonalChatController : ", message)

	id, err := manager.chatService.SavePersonalChat(senderEmailReq, receiverEmailReq, messageReq, message.CreatedAt)
	if err != nil {
		return
	}

	message.Id = id
	manager.SendMessageToUser(message)
}

func (manager *Manager) SendGroupChatController(messageRequest map[string]any) {
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

	groupIdReq, ok := messageRequest["group_id"].(float64)

	if !ok {
		log.Println("receiver_email error")
		return
	}

	message := web.ChatResponse{
		SenderEmail: senderEmailReq,
		GroupId:     int64(groupIdReq),
		Message:     messageReq,
		CreatedAt:   time.Now(),
	}

	id, err := manager.chatService.SaveGroupChat(senderEmailReq, int64(groupIdReq), messageReq, message.CreatedAt)
	if err != nil {
		return
	}

	message.Id = id
	manager.SendMessageToGroup(message)
}
