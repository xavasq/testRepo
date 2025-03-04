package service

import (
	"bluePlastic/internal/models"
	"bluePlastic/internal/repository"
	"context"
	"fmt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	if user.Name == "" {
		fmt.Errorf("имя пользователя не может быть пустым")
	}
	return s.repo.CreateUser(ctx, user)
}
