package entity

import (
	"test-kp-golang/src/domain/user/entity"
	"time"
)

type Transaction struct {
	ID                int
	UserID            int
	User              entity.User `gorm:"foreignKey:UserID"`
	AmountOtr         int
	AmountFee         int
	AmountInstallment int
	AmountInterest    int
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time `sql:"index"`
}
