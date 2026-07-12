package user

import (
	"context"
	"time"
)

type Repository interface {
	Create(ctx context.Context, u *User) error
	FindByID(ctx context.Context, id string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByGoogleID(ctx context.Context, googleID string) (*User, error)
	Update(ctx context.Context, u *User) error
	IncrementFailedLogin(ctx context.Context, id string) (int, error)
	ResetFailedLogin(ctx context.Context, id string) error
	SetLock(ctx context.Context, id string, until time.Time) error
}