package handler

import (
	"net/http"
	"strconv"
	"test-kp-golang/src/user/entity"
	"test-kp-golang/src/user/request"
	usecase "test-kp-golang/src/user/use-case"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	usecase *usecase.UserUsecase
}

func NewUserHandler(r *gin.Engine, usecase *usecase.UserUsecase) {
	handler := &UserHandler{usecase: usecase}

	r.POST("/users", handler.CreateUser)
	r.GET("/users/:id", handler.GetUserByID)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user request.RegisterUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bornDate, err := time.Parse("2006-01-02", user.BornDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := h.usecase.CreateUser(entity.User{
		FullName:          user.FullName,
		LegalName:         user.LegalName,
		Email:             user.Email,
		Password:          string(hashPassword),
		BornCity:          user.BornCity,
		BornDate:          bornDate,
		Income:            user.Income,
		SelfiePhotoPath:   user.SelfiePhotoPath,
		IdentityPhotoPath: user.IdentityPhotoPath,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.usecase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
