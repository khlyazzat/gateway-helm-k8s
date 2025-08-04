package profile

import (
	"context"
	"database/sql"
	"errors"
	"main/services/profile-service/internal/db/models"
	"main/services/profile-service/internal/db/postgres"
	"main/services/profile-service/internal/dto"

	profileRepo "main/services/profile-service/internal/db/repository/profile"
)

type Profile interface {
	GetProfile(ctx context.Context, userId int64, email string) (*dto.GetProfileResponse, error)
	UpdateProfile(ctx context.Context, userId int64, request *dto.UpdateProfileRequest) error
}

type profileService struct {
	profileRepo profileRepo.Profile
}

func (s *profileService) GetProfile(ctx context.Context, userId int64, email string) (*dto.GetProfileResponse, error) {
	profile, err := s.profileRepo.GetProfileByUserID(ctx, userId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) { //TODO
		return nil, err
	}
	if profile != nil {
		//return  //TODO
	}

	var profileResponse dto.GetProfileResponse

	profileResponse.Email = email

	if profile.FirstName != "" {
		profileResponse.FirstName = &profile.FirstName
	}
	if profile.LastName != "" {
		profileResponse.LastName = &profile.LastName
	}
	if profile.Age != 0 {
		profileResponse.Age = &profile.Age
	}
	if profile.Address != "" {
		profileResponse.Address = &profile.Address
	}
	if profile.Phone != "" {
		profileResponse.Phone = &profile.Phone
	}
	if profile.FirstName != "" {
		profileResponse.FirstName = &profile.FirstName
	}

	return &profileResponse, nil
}

func (s *profileService) UpdateProfile(ctx context.Context, userId int64, request *dto.UpdateProfileRequest) error {
	profile, err := s.profileRepo.GetProfileByUserID(ctx, userId)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if profile == nil {
		profile = new(models.Profile)
	}

	if request.FirstName != nil {
		profile.FirstName = *request.FirstName
	}
	if request.LastName != nil {
		profile.LastName = *request.LastName
	}
	if request.Age != nil {
		profile.Age = *request.Age
	}
	if request.Address != nil {
		profile.Address = *request.Address
	}
	if request.Phone != nil {
		profile.Phone = *request.Phone
	}
	if request.FirstName != nil {
		profile.FirstName = *request.FirstName
	}
	if errors.Is(err, sql.ErrNoRows) {
		profile = &models.Profile{
			UserID: userId,
		}
		err = s.profileRepo.CreateProfile(ctx, profile)
	} else {
		err = s.profileRepo.UpdateProfile(ctx, profile)
	}
	return nil
}

func NewProfileService(db postgres.DB) Profile {
	return &profileService{
		profileRepo: profileRepo.New(db),
	}
}
