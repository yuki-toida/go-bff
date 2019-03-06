package entity_email

import "go-bff/bff/domain/entities"

type Email struct {
	entities.Model
	Email  string
	UserID uint64
}

type email struct {
	repository Repository
}

func New(r Repository) *email {
	return &email{repository: r}
}

func (e *email) First(emailID uint64) (*Email, error) {
	return e.repository.First(emailID)
}

func (e *email) Create(userID uint64, emailAddr string) (*Email, error) {
	return e.repository.Create(userID, emailAddr)
}

func (e *email) Update(emailID uint64, emailAddr string) (*Email, error) {
	return e.repository.Update(emailID, emailAddr)
}
