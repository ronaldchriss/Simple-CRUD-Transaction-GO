package user

import "time"

type User struct {
	ID           int
	Name         string
	Username     string
	PasswordHash string
	Token        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
