package websocketProvider

import (
	"log"

	"github.com/MasLazu/CheatChatV2/model/web"
)

const (
	PersonalMessage = 1
)

func (manager *Manager) Router(messageReq web.WebsocketReqRes, sender *Client) {
	log.Println("Router : ", messageReq.MessageType)

	if messageMap, ok := messageReq.Body.(map[string]any); ok {
		switch messageReq.MessageType {
		case PersonalMessage:
			manager.SendPersonalChatController(messageMap, sender)
		}
	} else {
		log.Println("error while marshalling message")
	}
}
