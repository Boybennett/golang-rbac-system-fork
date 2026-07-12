package user

import "time"

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	ID               string
	FirstName		 string
	LastName		 string
	Email            string
	PasswordHash     string // empty if OAuth-only account
	Role             Role
	IsEmailVerified  bool
	AvatarUrl		 string
	FailedLoginCount int
	LockedUntil      *time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}