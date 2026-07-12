package user

import "time"

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	CreatedAt time.Time
	UpdatedAt time.Time

	ID           string
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
	AvatarUrl    string
	Role         Role

	LockedUntil *time.Time

	FailedLoginCount int
	IsEmailVerified  bool
}
