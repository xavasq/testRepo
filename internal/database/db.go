package database

import (
	"bluePlastic/internal/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

var DB *pgxpool.Pool

func ConnectDB() {
	cfg := config.LoadEnv()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("ошибка при подключении к базе данных")
	}
	DB = pool
	log.Println("база данных подключена успешно")
}
