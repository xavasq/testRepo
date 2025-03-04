package repository

import (
	"bluePlastic/internal/models"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (s *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := "INSERT INTO users(name, age) VALUES ($1,$2) RETURNING id"

	if err := s.db.QueryRow(ctx, query, user.Name, user.Age).Scan(&user.ID); err != nil {
		return err
	}
	return nil
}
