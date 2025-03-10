package entity

import "time"

type Product struct {
	ID          int
	Name        string
	AmountPrice int
	AmountShip  int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}
