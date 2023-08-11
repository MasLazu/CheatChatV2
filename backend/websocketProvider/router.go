package websocketProvider

import (
	"github.com/MasLazu/CheatChatV2/model/web"
	"log"
)

const (
	PersonalMessage = 1
)

func (manager *Manager) Router(messageReq web.WebsocketReqRes) {
	if messageMap, ok := messageReq.Body.(map[string]any); ok {
		switch messageReq.MessageType {
		case PersonalMessage:
			manager.SendPersonalChatController(messageMap)
		}
	} else {
		log.Println("error while marshalling message")
	}
}
