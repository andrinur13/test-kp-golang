package entity

import "time"

type UserTenor struct {
	ID           int
	UserID       int
	TenorInMonth int
	Amount       int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `sql:"index"`
}
