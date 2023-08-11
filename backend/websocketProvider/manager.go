package websocketProvider

import (
	"encoding/json"
	"github.com/MasLazu/CheatChatV2/model/web"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/MasLazu/CheatChatV2/helper"
	"github.com/MasLazu/CheatChatV2/repository"
	"github.com/MasLazu/CheatChatV2/service"
	"github.com/gorilla/websocket"
)

type ClientList struct {
	Clients map[*Client]bool
	sync.RWMutex
}

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			allowedOrigins := [...]string{os.Getenv("FRONTEND_DOMAIN")}
			requestOrigin := r.Header.Get("Origin")
			for _, origin := range allowedOrigins {
				if origin == requestOrigin {
					return true
				}
			}
			return false
		},
	}
)

type Manager struct {
	PersonalMessageType int
	GroupMessageType    int
	Clients             ClientList

	sessionService  service.SessionService
	chatService     service.ChatService
	groupRepository repository.GroupRepository
}

func NewManager(sessionService service.SessionService, chatService service.ChatService, groupRepository repository.GroupRepository) *Manager {
	return &Manager{
		PersonalMessageType: 1,
		GroupMessageType:    2,
		Clients: ClientList{
			Clients: make(map[*Client]bool),
		},
		sessionService:  sessionService,
		chatService:     chatService,
		groupRepository: groupRepository,
	}
}

func (manager *Manager) Connect(writer http.ResponseWriter, request *http.Request) {
	conn, err := websocketUpgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Println(err, " while upgrading to websocketProvider connection")
		helper.WriteResponse(writer, http.StatusBadRequest, "BAD REQUEST", web.MessageResponse{Message: "error while upgrading to websocketProvider connection"})
		return
	}

	user, err := manager.sessionService.Current(request, request.Context())
	if err != nil {
		log.Println(err)
		helper.WriteResponse(writer, http.StatusInternalServerError, "INTERNAL SERVER ERROR", err)
		return
	}

	groups, err := manager.groupRepository.GetUserGroupIds(request.Context(), user.Email)
	if err != nil {
		log.Println(err)
		helper.WriteResponse(writer, http.StatusInternalServerError, "INTERNAL SERVER ERROR", err)
		return
	}

	client := NewClient(conn, manager, user.Email, groups)
	manager.addClient(client)

	go client.ReadMessage()
	go client.WriteMessage()
}

func (manager *Manager) addClient(client *Client) {
	manager.Clients.Lock()
	defer manager.Clients.Unlock()
	manager.Clients.Clients[client] = true
}

func (manager *Manager) removeClient(client *Client) {
	manager.Clients.Lock()
	defer manager.Clients.Unlock()
	delete(manager.Clients.Clients, client)
}

func (manager *Manager) SendMessageToUser(recieiverEmail string, message web.ChatResponse) {
	for c := range manager.Clients.Clients {
		if c.UserEmail == recieiverEmail {
			response, err := json.Marshal(message)
			if err != nil {
				log.Println(err, " while marshaling message")
				return
			}
			if err := c.Conn.WriteMessage(websocket.TextMessage, response); err != nil {
				log.Println(err, " while sending message to ", recieiverEmail)
			}
		}
	}
}

//func (socket *WebsocketService) brodcastToGroupMember(message model.WebsocketRequest, sender *Client) {
//	chatsRepository := repository.NewChatsRepository()
//	newMessage := domain.Chat{
//		Sender:    sender.UserEmail,
//		Body:      message.Message,
//		CreatedAt: time.Now(),
//	}
//
//	if err := chatsRepository.PushNewChatToGroup(context.Background(), message.GroupID, newMessage); err != nil {
//		log.Println(err)
//		return
//	}
//
//	if err := chatsRepository.UpdateLastUpdateGroup(context.Background(), message.GroupID, newMessage.CreatedAt); err != nil {
//		log.Println(err)
//		return
//	}
//
//	for c := range socket.Clients.Clients {
//		for _, groupId := range c.GroupList {
//			if groupId == message.GroupID && c.UserId != sender.UserId {
//				response, err := json.Marshal(model.WbsocketResponse{
//					GroupID:   message.GroupID,
//					Sender:    sender.UserEmail,
//					Body:      message.Message,
//					CreatedAt: time.Now(),
//				})
//
//				if err != nil {
//					return
//				}
//
//				c.Conn.WriteMessage(websocketProvider.TextMessage, response)
//			}
//		}
//	}
//}
