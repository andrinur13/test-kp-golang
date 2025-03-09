package handler

import (
	"net/http"
	"test-kp-golang/src/domain/auth/request"
	"test-kp-golang/src/domain/auth/response"
	loginUseCase "test-kp-golang/src/domain/auth/use-case"
	"test-kp-golang/src/domain/user/entity"
	usecase "test-kp-golang/src/domain/user/use-case"
	"test-kp-golang/src/utils"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	usecase *usecase.UserUsecase
}

func NewAuthHandler(r *gin.Engine, usecase *usecase.UserUsecase) {
	handler := &AuthHandler{usecase: usecase}

	r.POST("/auth/login", handler.Login)
	r.POST("/auth/register", handler.Register)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var login request.LoginRequest
	if err := c.ShouldBindJSON(&login); err != nil {
		utils.FormatErrorResponse(c, "error", http.StatusBadRequest, err.Error(), nil)
		return
	}

	user, err := h.usecase.GetUserByEmail(login.Email)
	if err != nil {
		utils.FormatErrorResponse(c, "error", http.StatusBadRequest, err.Error(), nil)
		return
	}

	hashedPasswordStatus := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if hashedPasswordStatus != nil {
		utils.FormatErrorResponse(c, "error", http.StatusBadRequest, "invalid password", nil)
		return
	}

	token, err := loginUseCase.GenerateJWT(user.ID)
	if err != nil {
		utils.FormatErrorResponse(c, "error", http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.FormatResponse(c, "success", http.StatusOK, "login success", response.LoginResponse{
		Token: token,
	})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var register request.RegisterUserRequest
	if err := c.ShouldBindJSON(&register); err != nil {
		utils.FormatErrorResponse(c, "error", http.StatusBadRequest, err.Error(), nil)
		return
	}

	_, err := h.usecase.GetUserByEmail(register.Email)
	if err == nil {
		utils.FormatErrorResponse(c, "error", http.StatusBadRequest, "email already exists", nil)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.FormatErrorResponse(c, "error", http.StatusBadRequest, err.Error(), nil)
		return
	}

	bornDate, err := time.Parse("2006-01-02", register.BornDate)
	if err != nil {
		utils.FormatErrorResponse(c, "error", http.StatusBadRequest, "invalid date format", nil)
		return
	}

	_, err = h.usecase.CreateUser(entity.User{
		FullName:          register.FullName,
		LegalName:         register.LegalName,
		Email:             register.Email,
		Password:          string(hashedPassword),
		BornCity:          register.BornCity,
		BornDate:          bornDate,
		Income:            register.Income,
		IdentityPhotoPath: register.IdentityPhotoPath,
		SelfiePhotoPath:   register.SelfiePhotoPath,
	})

	if err != nil {
		utils.FormatErrorResponse(c, "error", http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.FormatResponse(c, "success", http.StatusOK, "register success", nil)
}
