package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mokhlesurr031/goblog/internal/models"
	"github.com/mokhlesurr031/goblog/internal/user/usecase"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(router *gin.Engine, userUsecase usecase.UserUsecase) {
	handler := &UserHandler{userUsecase: userUsecase}

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", handler.RegisterUser)
	}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userUsecase.RegisterUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
