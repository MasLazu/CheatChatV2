package websocket

import (
	"encoding/json"
	"github.com/MasLazu/CheatChatV2/helper"
	"github.com/MasLazu/CheatChatV2/model"
	"github.com/MasLazu/CheatChatV2/model/domain"
	"github.com/MasLazu/CheatChatV2/repository"
	"github.com/MasLazu/CheatChatV2/service"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
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
			allowedOrigins := [...]string{"http://localhost:5173"}
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

type WebsocketService struct {
	PersonalMessageType int
	GroupMessageType    int
	Clients             ClientList
}

func NewWebsocketService() *WebsocketService {
	return &WebsocketService{
		PersonalMessageType: 1,
		GroupMessageType:    2,
		Clients: ClientList{
			Clients: make(map[*Client]bool),
		},
	}
}

func (socket *WebsocketService) Connect(writer http.ResponseWriter, request *http.Request) {
	conn, err := websocketUpgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Println("upgrade error : ", err)
		helper.WriteResponse(writer, http.StatusBadRequest, "BAD REQUEST", model.MessageResponse{Message: "cannot upgrade to websocket connection"})
		return
	}

	sessionService := service.NewSessionService()
	user, err := sessionService.Current(request, request.Context())
	if err != nil {
		log.Println(err)
		helper.WriteResponse(writer, http.StatusInternalServerError, "INTERNAL SERVER ERROR", err)
		return
	}

	groupRepository := repository.NewGroupReposiroty()
	groups, err := groupRepository.GetAllUserGroupId(request.Context(), user.Id)
	log.Println("user_id : ", user.Id)
	log.Println(groups)
	if err != nil {
		log.Println(err)
		helper.WriteResponse(writer, http.StatusInternalServerError, "INTERNAL SERVER ERROR", err)
		return
	}

	client := NewClient(conn, socket, user.Email, user.Id, groups)
	socket.addClient(client)

	go client.ReadMessage()
	go client.WriteMessage()
}

func (socket *WebsocketService) sendMessageToUser(message model.WebsocketRequest, sender *Client) {
	chatsRepository := repository.NewChatsRepository()
	newMessage := domain.Chat{
		Sender:    sender.UserEmail,
		Body:      message.Message,
		CreatedAt: time.Now(),
	}

	if err := chatsRepository.PushNewChatToPersonalChat(context.Background(), sender.UserEmail, message.EmailDestination, newMessage); err != nil {
		log.Println(err)
		return
	}

	if err := chatsRepository.UpdateLastUpdatePersonalChat(context.Background(), sender.UserEmail, message.EmailDestination, newMessage.CreatedAt); err != nil {
		log.Println(err)
		return
	}

	for c := range socket.Clients.Clients {
		if c.UserEmail == message.EmailDestination {
			response, err := json.Marshal(model.WbsocketResponse{
				Sender:    sender.UserEmail,
				Body:      message.Message,
				CreatedAt: time.Now(),
			})

			if err != nil {
				return
			}
			c.Conn.WriteMessage(websocket.TextMessage, response)
		}
	}
}

func (socket *WebsocketService) brodcastToGroupMember(message model.WebsocketRequest, sender *Client) {
	chatsRepository := repository.NewChatsRepository()
	newMessage := domain.Chat{
		Sender:    sender.UserEmail,
		Body:      message.Message,
		CreatedAt: time.Now(),
	}

	if err := chatsRepository.PushNewChatToGroup(context.Background(), message.GroupID, newMessage); err != nil {
		log.Println(err)
		return
	}

	if err := chatsRepository.UpdateLastUpdateGroup(context.Background(), message.GroupID, newMessage.CreatedAt); err != nil {
		log.Println(err)
		return
	}

	for c := range socket.Clients.Clients {
		for _, groupId := range c.GroupList {
			if groupId == message.GroupID && c.UserId != sender.UserId {
				response, err := json.Marshal(model.WbsocketResponse{
					GroupID:   message.GroupID,
					Sender:    sender.UserEmail,
					Body:      message.Message,
					CreatedAt: time.Now(),
				})

				if err != nil {
					return
				}

				c.Conn.WriteMessage(websocket.TextMessage, response)
			}
		}
	}
}

func (socket *WebsocketService) addClient(client *Client) {
	socket.Clients.Lock()
	defer socket.Clients.Unlock()
	socket.Clients.Clients[client] = true
}

func (socket *WebsocketService) removeClient(client *Client) {
	socket.Clients.Lock()
	defer socket.Clients.Unlock()
	delete(socket.Clients.Clients, client)
}
