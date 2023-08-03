package domain

type Chat struct {
	Id          int64
	SenderEmail string
	Message     string
	ChatRoom    int64
}
