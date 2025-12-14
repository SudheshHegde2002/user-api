package service

import (
	"context"
	"time"
	"user-api/internal/models"
	"user-api/internal/repository"
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
) error {
	_, err := time.Parse("2006-01-02", dobStr)
	if err != nil {
		return err
	}

	_, err = s.repo.CreateUser(ctx, name, dobStr)
	return err
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
		Dob:  user.Dob.Time,
		Age:  age,
	}, nil
}

func (s *UserService) ListUsers(
	ctx context.Context,
) ([]models.UserResponse, error) {
	users, err := s.repo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]models.UserResponse, 0, len(users))

	for _, u := range users {
		resp = append(resp, models.UserResponse{
			ID:   u.ID,
			Name: u.Name,
			Dob:  u.Dob.Time,
			Age:  calculateAge(u.Dob.Time),
		})
	}

	return resp, nil
}
