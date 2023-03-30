package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Book struct {
	ID          int        `json:"id" gorm:"primaryKey;type:serial"`
	Title       string     `json:"title" gorm:"type:varchar(225)"`
	Author      string     `json:"author" gorm:"type:varchar(225)"`
	Description string     `json:"description" gorm:"type:text"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `sql:"default:null"`
}

func (b Book) Validation() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Title, validation.Required),
		validation.Field(&b.Author, validation.Required),
		validation.Field(&b.Description, validation.Required))
}
