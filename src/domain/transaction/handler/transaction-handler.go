package handler

import (
	"net/http"
	"strconv"
	"test-kp-golang/src/domain/transaction/request"
	usecase "test-kp-golang/src/domain/transaction/use-case"
	"test-kp-golang/src/domain/user/entity"
	"test-kp-golang/src/utils"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	useCase *usecase.TransactionUseCase
}

func NewTransactionHandler(r *gin.Engine, useCase *usecase.TransactionUseCase) {
	handler := &TransactionHandler{useCase: useCase}

	r.GET("/transactions", handler.FindByUserId)
	r.GET("/transactions/:id", handler.FindByUserId)
	r.POST("/transactions", handler.CreateTransaction)
}

func (h *TransactionHandler) FindByUserId(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(entity.User)
	transactions, err := h.useCase.FindByUserId(currentUser.ID)
	if err != nil {
		utils.FormatErrorResponse(c, "error", http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.FormatErrorResponse(c, "success", http.StatusOK, "success", transactions)
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(entity.User)

	var transaction request.TransactionRequest
	if err := c.ShouldBindJSON(&transaction); err != nil {
		utils.FormatErrorResponse(c, "error", http.StatusBadRequest, err.Error(), nil)
		return
	}

	transactionResponse, err := h.useCase.Create(currentUser.ID, transaction)
	if err != nil {
		utils.FormatErrorResponse(c, "error", http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.FormatErrorResponse(c, "success", http.StatusOK, "success", transactionResponse)
}

func (h *TransactionHandler) FindById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	transaction, err := h.useCase.FindById(id)
	if err != nil {
		utils.FormatErrorResponse(c, "error", http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.FormatErrorResponse(c, "success", http.StatusOK, "success", transaction)
}
