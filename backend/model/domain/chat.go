package domain

import "time"

type Chat struct {
	Id          int64
	SenderEmail string
	Message     string
	ChatRoom    int64
	CreatedAt   time.Time
}
