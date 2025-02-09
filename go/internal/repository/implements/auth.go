package implements

import (
	"context"
	"emregami/internal/domain"
	"emregami/internal/repository/repos"
	"errors"

	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) repos.AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) Save(ctx context.Context, auth *domain.Auth) error {
	return r.db.Create(auth).Error
}

func (r *authRepository) GetByUsername(ctx context.Context, username string) (*domain.Auth, error) {
	var auth domain.Auth
	err := r.db.Where("username = ?", username).First(&auth).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("failed to get user by username")
	}
	return &auth, nil
}

func (r *authRepository) GetByEmail(ctx context.Context, email string) (*domain.Auth, error) {
	var auth domain.Auth
	err := r.db.Where("email = ?", email).First(&auth).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("failed to get user by email")
	}
	return &auth, nil
}
