package handler

import (
	"net/http"
	"strconv"
	"test-kp-golang/src/domain/user/entity"
	"test-kp-golang/src/domain/user/request"
	"test-kp-golang/src/domain/user/response"
	usecase "test-kp-golang/src/domain/user/use-case"
	"test-kp-golang/src/utils"
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
		utils.FormatErrorResponse(c, "error", http.StatusBadRequest, err.Error(), nil)
		return
	}

	bornDate, err := time.Parse("2006-01-02", user.BornDate)
	if err != nil {
		utils.FormatErrorResponse(c, "error", http.StatusBadRequest, err.Error(), nil)
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		utils.FormatErrorResponse(c, "error", http.StatusBadRequest, err.Error(), nil)
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
		utils.FormatErrorResponse(c, "error", http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.FormatResponse(c, "success", http.StatusCreated, "User created successfully", response.UserRegistered{
		ID:       createdUser.ID,
		FullName: createdUser.FullName,
	})
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.FormatErrorResponse(c, "error", http.StatusBadRequest, err.Error(), nil)
		return
	}

	user, err := h.usecase.GetUserByID(id)
	if err != nil {
		utils.FormatErrorResponse(c, "error", http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.FormatResponse(c, "success", http.StatusOK, "User detail", response.UserDetail{
		ID:                user.ID,
		FullName:          user.FullName,
		LegalName:         user.LegalName,
		Email:             user.Email,
		BornCity:          user.BornCity,
		BornDate:          user.BornDate.Format("2006-01-02"),
		Income:            user.Income,
		IdentityPhotoPath: user.IdentityPhotoPath,
		SelfiePhotoPath:   user.SelfiePhotoPath,
		CreatedAt:         user.CreatedAt.Format("2006-01-02 15:04:05"),
	})
}
