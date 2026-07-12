package auth

import (
	"context"
)

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type Service interface {
	RegisterWithPassword(ctx context.Context, email, password string) error
	LoginWithPassword(ctx context.Context, email, password string) (*TokenPair, error)
	LoginWithGoogle(ctx context.Context, googleToken string) (*TokenPair, error)
	RequestMagicOTP(ctx context.Context, email string) error
	VerifyMagicOTP(ctx context.Context, email, otp string) (*TokenPair, error)
	VerifyEmail(ctx context.Context, token string) error
	RequestPasswordReset(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, token, newPassword string) error
	RefreshToken(ctx context.Context, refreshToken string) (*TokenPair, error)
	Logout(ctx context.Context, refreshToken string) error
}

type authService struct {
	userRepository Repository
}

func NewService(userRepository Repository) Service {
	return &authService{
		userRepository: userRepository,
	}
}

func (s *authService) RegisterWithPassword(ctx context.Context, email, password string) error            {
	return nil
}
func (s *authService) LoginWithPassword(ctx context.Context, email, password string) (*TokenPair, error) {
	return nil, nil
}
func (s *authService) LoginWithGoogle(ctx context.Context, googleToken string) (*TokenPair, error)       {
	return nil, nil
}
func (s *authService) RequestMagicOTP(ctx context.Context, email string) error                           {
	return nil
}
func (s *authService) VerifyMagicOTP(ctx context.Context, email, otp string) (*TokenPair, error)         {
	return nil, nil
}
func (s *authService) VerifyEmail(ctx context.Context, token string) error                               {
	return nil
}
func (s *authService) RequestPasswordReset(ctx context.Context, email string) error                      {
	return nil
}
func (s *authService) ResetPassword(ctx context.Context, token, newPassword string) error                {
	return nil
}
func (s *authService) RefreshToken(ctx context.Context, refreshToken string) (*TokenPair, error)         {
	return nil, nil 
}
func (s *authService) Logout(ctx context.Context, refreshToken string) error                             {
	return nil
}
