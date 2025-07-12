package cache

import (
	"context"
	"time"
)

type TokenCache interface {
	SetRefreshToken(ctx context.Context, userID string, refreshToken string) error
	RefreshTokenExist(ctx context.Context, userID string, refreshToken string) (bool, error)
}

type tokenCache struct {
	Cache
	RefreshTTL int
}

func (r *tokenCache) SetRefreshToken(ctx context.Context, userID string, refreshToken string) error {
	return r.Set(ctx, userID, refreshToken, time.Duration(r.RefreshTTL)*time.Minute)
}

func (r *tokenCache) RefreshTokenExist(ctx context.Context, userID string, refreshToken string) (bool, error) {
	exists, err := r.Exists(ctx, userID)
	if !exists {
		return false, err
	}
	token, err := r.Get(ctx, userID)
	if err != nil {
		return false, err
	}
	return refreshToken == token, nil
}

func NewTokenCache(client Cache, refreshTTL int) TokenCache {
	return &tokenCache{
		RefreshTTL: refreshTTL,
		Cache:      client,
	}
}
