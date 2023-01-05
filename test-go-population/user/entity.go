package user

import "time"

type User struct {
	ID          int
	FullName    string
	UserName    string
	Email       string
	Password    string
	ProfilePict string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
