package interfaces

import "emregami/internal/domain"

type AuthRepository interface {
	Create(auth *domain.Auth) (*domain.Auth, error)
	GetByEmail(email string) (*domain.Auth, error)
	GetByUsername(username string) (*domain.Auth, error)
}
