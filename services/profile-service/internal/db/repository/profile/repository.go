package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"main/services/profile-service/internal/db/models"
	"main/services/profile-service/internal/db/repository"
	"main/services/profile-service/internal/values"

	db "main/services/profile-service/internal/db/postgres"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

type Profile interface {
	CreateProfile(ctx context.Context, profile *models.Profile) error
	GetProfileByUserID(ctx context.Context, userId int64) (*models.Profile, error)
	UpdateProfile(ctx context.Context, profile *models.Profile) error
}

type profileRepository struct {
	repository.CRUD
	db.DB
}

func (r *profileRepository) CreateProfile(ctx context.Context, profile *models.Profile) error {
	err := r.DB.NewInsert().
		Model(profile).
		Returning("id").
		Scan(ctx)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return fmt.Errorf("%w: %s", values.ErrEmailExists, pgErr.Detail)
		}
		return fmt.Errorf("failed to insert profile: %w", err)
	}
	return nil
}

func (r *profileRepository) GetProfileByUserID(ctx context.Context, userId int64) (*models.Profile, error) {
	profile := new(models.Profile)

	err := r.DB.NewSelect().
		Model(profile).
		Where("user_id = ?", userId).
		Limit(1).
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
		// return nil, fmt.Errorf("%w: id=%d", values.ErrUserNotFound, userId)
	}
	if err != nil {
		return nil, err
		// return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	return profile, nil
}

func (r *profileRepository) UpdateProfile(ctx context.Context, profile *models.Profile) error {
	err := r.DB.NewUpdate().
		Model(profile).
		WherePK().
		Returning("*").
		Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("%w: id=%d", values.ErrUserNotFound, profile.ID)
	}
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func New(db db.DB) Profile {
	return &profileRepository{
		CRUD: repository.NewCRUD(db),
		DB:   db,
	}
}
