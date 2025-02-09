package configs

import (
	"emregami/pkg/utils"
	"time"
)

type TokenConfig struct {
	AccessTokenSecret  string        `mapstructure:"access_token_secret"`
	RefreshTokenSecret string        `mapstructure:"refresh_token_secret"`
	AccessTokenExpiry  time.Duration `mapstructure:"access_token_expiry"`
	RefreshTokenExpiry time.Duration `mapstructure:"refresh_token_expiry"`
}

func NewTokenConfig() *TokenConfig {
	return &TokenConfig{
		AccessTokenSecret:  utils.GetString("ACCESS_TOKEN_SECRET"),
		RefreshTokenSecret: utils.GetString("REFRESH_TOKEN_SECRET"),
		AccessTokenExpiry:  utils.GetDuration("ACCESS_TOKEN_EXPIRY"),
		RefreshTokenExpiry: utils.GetDuration("REFRESH_TOKEN_EXPIRY"),
	}
}
