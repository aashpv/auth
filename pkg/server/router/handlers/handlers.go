package handlers

import (
	"errors"
	"github.com/aashpv/auth/pkg/models"
	"github.com/aashpv/auth/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handlers interface {
	SignUp(c *gin.Context)
	Login(c *gin.Context)
}

type handlers struct {
	s service.Service
}

func New(src service.Service) Handlers {
	return &handlers{s: src}
}

func (h *handlers) SignUp(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err := h.s.SignUp(user)
	if err != nil {
		if errors.Is(err, errors.New("invalid email")) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid email"})
			return
		}
		if errors.Is(err, errors.New("invalid password")) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
			return
		}
		if errors.Is(err, errors.New("invalid phone number")) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid phone number"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created successfully"})
}

func (h *handlers) Login(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	token, err := h.s.SignIn(user.Email, user.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
		return
	}

	c.JSON(http.StatusOK, token)
}
