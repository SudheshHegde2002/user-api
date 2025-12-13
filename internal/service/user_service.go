package service

import (
	"context"
	"time"
	"user-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

