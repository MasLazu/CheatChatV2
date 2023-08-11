package websocketProvider

import (
	"encoding/json"
	"log"
	"time"

	"github.com/MasLazu/CheatChatV2/model"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn        *websocket.Conn
	UserEmail   string
	GroupList   []int64
	MessageChan chan model.WebsocketReqRes
	Manager     *Manager
}

const (
	pongWait     = 20 * time.Second
	pingInterval = pongWait * 9 / 10
)

func NewClient(conn *websocket.Conn, manager *Manager, email string, groupList []int64) *Client {
	return &Client{
		Conn:        conn,
		UserEmail:   email,
		GroupList:   groupList,
		MessageChan: make(chan model.WebsocketReqRes),
		Manager:     manager,
	}
}

func (client *Client) PongHandler(pongMessage string) error {
	log.Println(pongMessage, " from ", client.UserEmail)
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
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println(err, " reading message from ", client.UserEmail)
			}
			break
		}

		var request model.WebsocketReqRes
		if err := json.Unmarshal(payload, &request); err != nil {
			log.Println(err, " marshalling message from ", client.UserEmail)
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
			client.Manager.Router(message)
		case <-ticker.C:
			log.Println("current client : ", len(client.Manager.Clients.Clients))
			log.Println("sending ping to ", client.UserEmail)
			if err := client.Conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Println(err, "while sending ping to ", client.UserEmail)
				return
			}
		}
	}
}
