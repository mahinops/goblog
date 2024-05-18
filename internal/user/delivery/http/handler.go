package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mokhlesurr031/goblog/internal/middleware"
	"github.com/mokhlesurr031/goblog/internal/models"
	"github.com/mokhlesurr031/goblog/internal/user/usecase"
	"github.com/mokhlesurr031/goblog/pkg/jwt"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(router *gin.Engine, userUsecase usecase.UserUsecase) {
	handler := &UserHandler{userUsecase: userUsecase}

	userRoutes := router.Group("/users")
	{
		// Routes excluded from authentication middleware
		userRoutes.POST("/register", handler.RegisterUser)
		userRoutes.POST("/login", handler.Login)

		// Routes that require authentication middleware
		userRoutes.GET("/me", middleware.AuthMiddleware(), handler.GetCurrentUser)
	}
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
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

func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userUsecase.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token, "user": user})
}

func (h *UserHandler) GetCurrentUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	user, err := h.userUsecase.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch user info"})
		return
	}

	c.JSON(http.StatusOK, user)
}
