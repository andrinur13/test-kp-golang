package usecase

import (
	"errors"
	productRepository "test-kp-golang/src/domain/product/repository"
	transactionEntity "test-kp-golang/src/domain/transaction/entity"
	"test-kp-golang/src/domain/transaction/repository"
	"test-kp-golang/src/domain/transaction/request"
	"test-kp-golang/src/domain/transaction/response"
	tenorEntity "test-kp-golang/src/domain/user-tenor/entity"
	tenorRepository "test-kp-golang/src/domain/user-tenor/repository"
	"time"
)

type TransactionUseCase struct {
	transactionRepository repository.TransactionRepository
	tenorRepository       tenorRepository.UserTenorRepository
	productRepository     productRepository.ProductRepository
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

func (u *TransactionUseCase) Create(userId int, request request.TransactionRequest) (response.TransactionResponse, error) {
	userTenor, err := u.tenorRepository.FindByUserId(userId)

	if err != nil {
		return response.TransactionResponse{}, err
	}

	var selectedTenor tenorEntity.UserTenor
	for _, tenor := range userTenor {
		if tenor.TenorInMonth == request.Tenor {
			selectedTenor = tenor
			break
		}
	}

	unpaidTransactions, err := u.transactionRepository.FindUnpaidStatusTransactionByUserId(userId)

	if err != nil {
		return response.TransactionResponse{}, err
	}

	amountUnpaid := 0
	for _, unpaidTransaction := range unpaidTransactions {
		amountUnpaid += unpaidTransaction.AmountOtr + unpaidTransaction.AmountFee + unpaidTransaction.AmountInterest
	}

	if amountUnpaid > selectedTenor.Amount {
		return response.TransactionResponse{}, errors.New("Your limit is not enough")
	}

	product, err := u.productRepository.GetProductByID(request.ProducId)
	if err != nil {
		return response.TransactionResponse{}, err
	}

	amountFeePct := 10
	interestPc := 5

	amountFee := (amountFeePct / 100) * product.AmountPrice
	amountInstallment := selectedTenor.TenorInMonth
	amountInterest := (interestPc / 100) * product.AmountPrice

	transaction := transactionEntity.Transaction{
		UserID:            userId,
		AmountOtr:         product.AmountPrice,
		AmountFee:         amountFee,
		AmountInstallment: amountInstallment,
		AmountInterest:    amountInterest,
	}

	createdTransaction, err := u.transactionRepository.Create(transaction)

	if err != nil {
		return response.TransactionResponse{}, err
	}

	transactionResponse := response.TransactionResponse{
		ID:                createdTransaction.ID,
		AmountOtr:         createdTransaction.AmountOtr,
		AmountFee:         createdTransaction.AmountFee,
		AmountInstallment: createdTransaction.AmountInstallment,
		AmountInterest:    createdTransaction.AmountInterest,
		CreatedAt:         createdTransaction.CreatedAt.Format(time.RFC3339),
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
