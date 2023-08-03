package domain

type Group struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	ChatRoom int64  `json:"chat_room,omitempty"`
}
