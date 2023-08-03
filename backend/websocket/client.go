package websocket

import (
	"encoding/json"
	"log"
	"runtime"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn        *websocket.Conn
	UserEmail   string
	GroupList   []int32
	MessageChan chan interface{}
	Manager     *WebsocketService
}

const (
	pongWait     = 20 * time.Second
	pingInterval = pongWait * 9 / 10
)

func NewClient(conn *websocket.Conn, manager *WebsocketService, email string, userId int32, groupList []int32) *Client {
	return &Client{
		Conn:        conn,
		UserEmail:   email,
		GroupList:   groupList,
		MessageChan: make(chan interface{}),
		Manager:     manager,
	}
}

func (client *Client) PongHandler(pongMessage string) error {
	log.Println("pong : ", client.UserEmail)
	return client.Conn.SetReadDeadline(time.Now().Add(pongWait))
}

func (client *Client) ReadMessage() {
	defer func() {
		client.Manager.removeClient(client)
	}()

	if err := client.Conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println(err)
		log.Println(client, err)
		return
	}

	client.Conn.SetPongHandler(client.PongHandler)

	for {
		_, payload, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v", err)
			}
			break
		}

		var request interface{}
		if err := json.Unmarshal(payload, &request); err != nil {
			log.Printf("error marshalling message: %v", err)
			break
		}

		client.MessageChan <- request
	}
}

func (client *Client) WriteMessage() {
	ticker := time.NewTicker(pingInterval)

	defer func() {
		ticker.Stop()
		client.Manager.removeClient(client)
	}()

	for {
		select {
		case message := <-client.MessageChan:
			// disini harusnya websocket router(message, client)
			//if message.DestinationType == client.Manager.PersonalMessageType {
			//	client.Manager.sendMessageToUser(message, client)
			//} else if message.DestinationType == client.Manager.GroupMessageType {
			//	client.Manager.brodcastToGroupMember(message, client)
			//}
		case <-ticker.C:
			log.Println("cuncurent client : ", len(client.Manager.Clients.Clients))
			log.Println("ping : ", client, "goroutine : ", runtime.NumGoroutine())
			if err := client.Conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Println("write message: ", err)
				return
			}
		}
	}
}
