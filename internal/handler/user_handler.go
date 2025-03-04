package handler

import (
	"bluePlastic/internal/models"
	"bluePlastic/internal/service"
	"context"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(404, gin.H{"error": "неккоректные данные"})
		return
	}

	if err := h.service.CreateUser(context.Background(), &user); err != nil {
		c.JSON(404, gin.H{"error": "ошибка при создании пользователя"})
		return
	}
	c.JSON(201, gin.H{"user": user})
}
