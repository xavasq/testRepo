package router

import (
	"bluePlastic/internal/handler"
	"bluePlastic/internal/repository"
	"bluePlastic/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(r *gin.Engine, db *pgxpool.Pool) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	api := r.Group("/api")
	{
		api.POST("/users", userHandler.CreateUser)
		api.GET("/users/:name", userHandler.GetUserByEmail)
	}
}
