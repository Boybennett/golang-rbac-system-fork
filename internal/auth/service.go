package auth

import (
	"context"
	"errors"
	"log"
	"github.com/Steve-s-Circle-on-System-Design/golang-rbac-system/internal/user"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserWithEmailAlreadyExists = errors.New("user with that email already exists")

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type Service interface {
	RegisterWithPassword(ctx context.Context, email, password string) error
}

type authService struct {
	userRepository user.Repository
}

func NewService(userRepository user.Repository) Service {
	return &authService{
		userRepository: userRepository,
	}
}

func (s *authService) RegisterWithPassword(ctx context.Context, email, password string) error {
	existingUser, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Println("failed to check existing user:", err)
			return err
		}
		existingUser = nil
	}

	if existingUser != nil {
		return ErrUserWithEmailAlreadyExists
	}
	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Something went wrong while hashing the password", err.Error())
		return err
	}
	newUser := &user.User{
		Email:        email,
		PasswordHash: string(passwordHash),
	}
	err = s.userRepository.Create(ctx, newUser)
	if err != nil {
		log.Println("Something went wrong while trying to save the new user in the db", err.Error())
	}
	return nil
}
