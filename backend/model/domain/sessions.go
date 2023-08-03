package domain

import "time"

type Session struct {
	UserEmail string
	Token     string
	ExpiredAt time.Time
}
