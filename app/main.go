package main

import (
	"bluePlastic/internal/database"
	"bluePlastic/internal/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	database.ConnectDB()
	defer database.DB.Close()

	r := gin.Default()

	router.SetupRouter(r, database.DB)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("ошибка при запуске сервера")
	}
}
