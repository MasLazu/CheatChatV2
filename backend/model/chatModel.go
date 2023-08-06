package model

import "time"

type PreviewGroupChat struct {
	GroupId     int64     `json:"group_id"`
	GroupName   string    `json:"group_name"`
	ChatId      int64     `json:"chat_id"`
	SenderEmail string    `json:"sender_email"`
	Message     string    `json:"message"`
	CreatedAt   time.Time `json:"created_at"`
}

type PreviewPersonalChat struct {
	Email       string    `json:"email"`
	ChatId      int64     `json:"chat_id"`
	SenderEmail string    `json:"sender_email"`
	Message     string    `json:"message"`
	CreatedAt   time.Time `json:"created_at"`
}

type ChatResponse struct {
	Id          int64     `json:"id"`
	SenderEmail string    `json:"sender_email"`
	Message     string    `json:"message"`
	CreatedAt   time.Time `json:"created_at"`
}

type PreviewChatResponse struct {
	Group    []PreviewGroupChat    `json:"group"`
	Personal []PreviewPersonalChat `json:"personal"`
}
