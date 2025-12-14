package service

import (
	"context"
	"time"
	"user-api/internal/logger"
	"user-api/internal/models"
	"user-api/internal/repository"

	"go.uber.org/zap"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(
	ctx context.Context,
	name string,
	dobStr string,
) (*models.UserResponse, error) {
	_, err := time.Parse("2006-01-02", dobStr)
	if err != nil {
		logger.Log.Warn("invalid dob format",
			zap.String("dob", dobStr),
		)
		return nil, err
	}

	u, err := s.repo.CreateUser(ctx, name, dobStr)
	if err != nil {
		return nil, err
	}
	return &models.UserResponse{
		ID:   u.ID,
		Name: u.Name,
		Dob:  u.Dob.Time.Format("2006-01-02"),
	}, nil
}

func (s *UserService) GetUserByID(
	ctx context.Context,
	id int32,
) (*models.UserResponse, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	age := calculateAge(user.Dob.Time)

	return &models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Time.Format("2006-01-02"),
		Age:  age,
	}, nil
}

func (s *UserService) ListUsers(
	ctx context.Context,
	limit int32,
	offset int32,
) ([]models.UserResponse, error) {

	users, err := s.repo.ListUsers(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	resp := make([]models.UserResponse, 0, len(users))

	for _, u := range users {
		resp = append(resp, models.UserResponse{
			ID:   u.ID,
			Name: u.Name,
			Dob:  u.Dob.Time.Format("2006-01-02"),
			Age:  calculateAge(u.Dob.Time),
		})
	}

	return resp, nil
}

func (s *UserService) UpdateUser(
	ctx context.Context,
	id int32,
	name string,
	dobStr string,
) (*models.UserResponse, error) {

	if _, err := time.Parse("2006-01-02", dobStr); err != nil {
		logger.Log.Warn("invalid dob format for update",
			zap.Int32("user_id", id),
			zap.String("dob", dobStr),
		)
		return nil, err
	}

	u, err := s.repo.UpdateUser(ctx, id, name, dobStr)
	if err != nil {
		return nil, err
	}

	return &models.UserResponse{
		ID:   u.ID,
		Name: u.Name,
		Dob:  u.Dob.Time.Format("2006-01-02"),
	}, nil
}

func (s *UserService) DeleteUser(
	ctx context.Context,
	id int32,
) error {
	return s.repo.DeleteUser(ctx, id)
}
