package entity

import (
	productEntity "test-kp-golang/src/domain/product/entity"
	"test-kp-golang/src/domain/user/entity"
	"time"
)

type Transaction struct {
	ID                int
	UserID            int
	ProductID         int
	User              entity.User `gorm:"foreignKey:UserID"`
	Product           productEntity.Product
	AmountOtr         int
	AmountFee         int
	AmountInstallment int
	AmountInterest    int
	Status            string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time `sql:"index"`
}
