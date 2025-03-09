package entity

import "time"

type User struct {
	ID                int
	FullName          string
	LegalName         string
	Email             string
	Password          string
	BornCity          string
	BornDate          time.Time `sql:"type:date"`
	Income            int
	IdentityPhotoPath string
	SelfiePhotoPath   string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time `sql:"index"`
}
