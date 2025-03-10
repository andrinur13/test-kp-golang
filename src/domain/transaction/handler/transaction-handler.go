package handler

import (
	"test-kp-golang/src/domain/transaction/response"
	usecase "test-kp-golang/src/domain/transaction/use-case"
)

type TransactionHandler struct {
	useCase *usecase.TransactionUseCase
}

func NewTransactionHandler(useCase *usecase.TransactionUseCase) *TransactionHandler {
	return &TransactionHandler{
		useCase: useCase,
	}
}

func (h *TransactionHandler) FindByUserId(id int) ([]response.TransactionResponse, error) {
	transactions, err := h.useCase.FindByUserId(id)
	if err != nil {
		return []response.TransactionResponse{}, err
	}

	return transactions, nil
}

// func (h *TransactionHandler) Create(transactionRequest request.TransactionRequest) (response.TransactionResponse, error) {

// 	transaction, err := h.useCase.Create(transactionRequest)
// 	if err != nil {
// 		return response.TransactionResponse{}, err
// 	}

// 	return transaction, nil
// }

func (h *TransactionHandler) FindById(id int) (response.TransactionResponse, error) {
	transaction, err := h.useCase.FindById(id)
	if err != nil {
		return response.TransactionResponse{}, err
	}

	return transaction, nil
}
