package repos

import (
	"context"
	"emregami/internal/domain"
)

type AuthRepository interface {
	Save(ctx context.Context, auth *domain.Auth) error
	GetByEmail(ctx context.Context, email string) (*domain.Auth, error)
	GetByUsername(ctx context.Context, username string) (*domain.Auth, error)
}
