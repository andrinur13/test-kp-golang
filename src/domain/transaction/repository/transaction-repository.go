package repository

import (
	"test-kp-golang/src/domain/transaction/entity"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (r *TransactionRepository) FindByUserId(id int) ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	result := r.db.Where("user_id = ?", id).Find(&transactions)
	if result.Error != nil {
		return transactions, result.Error
	}

	return transactions, nil
}

func (r *TransactionRepository) Create(transaction entity.Transaction) (entity.Transaction, error) {
	result := r.db.Create(&transaction)
	if result.Error != nil {
		return transaction, result.Error
	}

	return transaction, nil
}

func (r *TransactionRepository) FindById(id int) (entity.Transaction, error) {
	var transaction entity.Transaction

	result := r.db.Where("id = ?", id).First(&transaction)
	if result.Error != nil {
		return transaction, result.Error
	}

	return transaction, nil
}

func (r *TransactionRepository) FindPaidStatusTransactionByUserId(userId int) ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	result := r.db.Where("user_id = ? AND status = ?", userId, "PAID").Find(&transactions)
	if result.Error != nil {
		return transactions, result.Error
	}

	return transactions, nil
}

func (r *TransactionRepository) FindUnpaidStatusTransactionByUserId(userId int) ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	result := r.db.Where("user_id = ? AND status = ?", userId, "UNPAID").Find(&transactions)
	if result.Error != nil {
		return transactions, result.Error
	}

	return transactions, nil
}
