package auth

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	"main/services/auth-service/internal/db/cache"
	"main/services/auth-service/internal/db/models"
	"main/services/auth-service/internal/db/postgres"
	userRepo "main/services/auth-service/internal/db/repository/user"
	"main/services/auth-service/internal/values"
	"main/services/auth-service/jwt"
	"main/services/auth-service/utils"

	"main/services/auth-service/internal/dto"
)

type Auth interface {
	SignUp(ctx context.Context, request *dto.SignUpRequest) (*dto.SignUpResponse, error)
	SignIn(ctx context.Context, request *dto.SignInRequest) (*dto.TokenResponse, error)
	Refresh(ctx context.Context, request *dto.RefreshRequest) (*dto.TokenResponse, error)
}

type authService struct {
	userRepo   userRepo.User
	tokenCache cache.TokenCache
	jwt        jwt.JWT
}

func (s *authService) SignUp(ctx context.Context, request *dto.SignUpRequest) (*dto.SignUpResponse, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, request.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if user != nil {
		return nil, values.ErrEmailExists
	}

	newUser := &models.User{
		Name:      request.Name,
		Password:  utils.HashPassword(request.Password),
		Email:     request.Email,
		Age:       request.Age,
		CreatedAt: time.Now(),
	}

	_, err = s.userRepo.AddUser(ctx, newUser)
	if err != nil {
		return nil, err
	}
	return &dto.SignUpResponse{}, nil //TODO
}

func (s *authService) SignIn(ctx context.Context, request *dto.SignInRequest) (*dto.TokenResponse, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, request.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if user != nil {
		return nil, values.ErrEmailExists
	}
	if !utils.CheckPassword(request.Password, user.Password) {
		return nil, values.ErrWrongLoginOrPassword
	}
	jwtToken, ttl, err := s.jwt.GenerateToken(strconv.Itoa(int(user.ID)), jwt.Access)
	if err != nil {
		return nil, err
	}
	refreshToken, _, err := s.jwt.GenerateToken(strconv.Itoa(int(user.ID)), jwt.Refresh)
	if err != nil {
		return nil, err
	}
	if err = s.tokenCache.SetRefreshToken(ctx, strconv.Itoa(int(user.ID)), refreshToken); err != nil {
		return nil, values.ErrSetRefreshToken
	}
	return &dto.TokenResponse{
		AuthToken:    jwtToken,
		RefreshToken: refreshToken,
		TTL:          ttl,
	}, nil
}

func (s *authService) Refresh(ctx context.Context, request *dto.RefreshRequest) (*dto.TokenResponse, error) {
	claims, err := s.jwt.ParseJWTToken(request.Token)
	if err != nil {
		return nil, values.ErrParseToken
	}
	exists, err := s.tokenCache.RefreshTokenExist(ctx, claims.Id, request.Token)
	if err != nil || !exists {
		return nil, values.ErrGetRefreshToken
	}
	jwtToken, ttl, err := s.jwt.GenerateToken(claims.Id, jwt.Access)
	if err != nil {
		return nil, err
	}
	refreshToken, _, err := s.jwt.GenerateToken(claims.Id, jwt.Refresh)
	if err != nil {
		return nil, err
	}
	if err = s.tokenCache.SetRefreshToken(ctx, claims.Id, refreshToken); err != nil {
		return nil, values.ErrSetRefreshToken
	}
	return &dto.TokenResponse{
		AuthToken:    jwtToken,
		RefreshToken: refreshToken,
		TTL:          ttl,
	}, nil
}

func NewAuthService(db postgres.DB, tokenCache cache.TokenCache, jwt jwt.JWT) Auth {
	return &authService{
		userRepo:   userRepo.New(db),
		tokenCache: tokenCache,
		jwt:        jwt,
	}
}
