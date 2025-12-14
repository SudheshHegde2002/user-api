package service

import (
	"context"
	"time"
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
)error{
	_, err := time.Parse("2006-01-02", dobStr)
	if err != nil {
		return err
	}

	_, err = s.repo.CreateUser(ctx, name, dobStr)
	return err
}
