package domain

import "time"

type Chat struct {
	Id          int64     `json:"id,omitempty"`
	SenderEmail string    `json:"sender_email,omitempty"`
	Message     string    `json:"message,omitempty"`
	ChatRoom    int64     `json:"chat_room,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}
