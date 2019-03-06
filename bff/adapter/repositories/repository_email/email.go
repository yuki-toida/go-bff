package repository_email

import (
	"github.com/jinzhu/gorm"
	"go-bff/bff/domain/entities"
	"go-bff/bff/domain/entities/entity_email"
)

type Email struct {
	entities.Model
	Email  string `gorm:"not null"`
	UserID uint64 `gorm:"not null;index"`
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *repository {
	return &repository{db: db}
}

func ToEntity(email *Email) *entity_email.Email {
	return &entity_email.Email{
		Model:  email.Model,
		Email:  email.Email,
		UserID: email.UserID,
	}
}

func (r *repository) First(emailID uint64) (*entity_email.Email, error) {
	email := &Email{Model: entities.Model{ID: emailID}}
	if err := r.db.First(email).Error; err != nil {
		return nil, err
	}
	return ToEntity(email), nil
}

func (r *repository) Create(userID uint64, emailAddr string) (*entity_email.Email, error) {
	e := Email{Email: emailAddr, UserID: userID}
	if err := r.db.Create(&e).Error; err != nil {
		return nil, err
	}
	return ToEntity(&e), nil
}

func (r *repository) Update(emailID uint64, emailAddr string) (*entity_email.Email, error) {
	e, err := r.First(emailID)
	if err != nil {
		return nil, err
	}
	e.Email = emailAddr
	if err := r.db.Save(e).Error; err != nil {
		return nil, err
	}
	return e, nil
}
