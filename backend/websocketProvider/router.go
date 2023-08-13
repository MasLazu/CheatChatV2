package websocketProvider

import (
	"log"

	"github.com/MasLazu/CheatChatV2/model/web"
)

const (
	PersonalMessage = 1
	GroupMessage    = 2
)

func (manager *Manager) Router(messageReq web.WebsocketReqRes) {
	log.Println("Router : ", messageReq.MessageType)

	if messageMap, ok := messageReq.Body.(map[string]any); ok {
		switch messageReq.MessageType {

		case PersonalMessage:
			manager.SendPersonalChatController(messageMap)

		case GroupMessage:
			manager.SendGroupChatController(messageMap)

		}
	} else {
		log.Println("error while marshalling message")
	}
}
