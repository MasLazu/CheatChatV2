package domain

import "time"

type User struct {
	Email     string
	Username  string
	Password  string
	CreatedAt time.Time
}
