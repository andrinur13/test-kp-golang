package usecase

import (
	"test-kp-golang/src/domain/transaction/entity"
	"test-kp-golang/src/domain/transaction/repository"
	"test-kp-golang/src/domain/transaction/response"
	"time"
)

type TransactionUseCase struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionUseCase(transactionRepository repository.TransactionRepository) *TransactionUseCase {
	return &TransactionUseCase{
		transactionRepository: transactionRepository,
	}
}

func (u *TransactionUseCase) FindByUserId(id int) ([]response.TransactionResponse, error) {
	transactions, err := u.transactionRepository.FindByUserId(id)
	if err != nil {
		return []response.TransactionResponse{}, err
	}

	var transactionResponses []response.TransactionResponse
	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, response.TransactionResponse{
			ID:                transaction.ID,
			AmountOtr:         transaction.AmountOtr,
			AmountFee:         transaction.AmountFee,
			AmountInstallment: transaction.AmountInstallment,
			AmountInterest:    transaction.AmountInterest,
			CreatedAt:         transaction.CreatedAt.Format(time.RFC3339),
		})
	}

	return transactionResponses, nil
}

func (u *TransactionUseCase) Create(transaction entity.Transaction) (response.TransactionResponse, error) {
	transaction, err := u.transactionRepository.Create(transaction)
	if err != nil {
		return response.TransactionResponse{}, err
	}

	transactionResponse := response.TransactionResponse{
		ID:                transaction.ID,
		AmountOtr:         transaction.AmountOtr,
		AmountFee:         transaction.AmountFee,
		AmountInstallment: transaction.AmountInstallment,
		AmountInterest:    transaction.AmountInterest,
		CreatedAt:         transaction.CreatedAt.Format(time.RFC3339),
	}

	return transactionResponse, nil
}

func (u *TransactionUseCase) FindById(id int) (response.TransactionResponse, error) {
	transaction, err := u.transactionRepository.FindById(id)
	if err != nil {
		return response.TransactionResponse{}, err
	}

	transactionResponse := response.TransactionResponse{
		ID:                transaction.ID,
		AmountOtr:         transaction.AmountOtr,
		AmountFee:         transaction.AmountFee,
		AmountInstallment: transaction.AmountInstallment,
		AmountInterest:    transaction.AmountInterest,
		CreatedAt:         transaction.CreatedAt.Format(time.RFC3339),
	}

	return transactionResponse, nil
}
