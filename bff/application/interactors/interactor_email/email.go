package interactor_email

import (
	"go-bff/bff/domain/entities/entity_email"
)

type interactor struct {
	emailRepository entity_email.Repository
}

func New(er entity_email.Repository) *interactor {
	return &interactor{emailRepository: er}
}

func (i *interactor) First(emailID uint64) (*entity_email.Email, error) {
	ee := entity_email.New(i.emailRepository)
	email, err := ee.First(emailID)
	if err != nil {
		return nil, err
	}
	return email, nil
}

func (i *interactor) Update(emailID uint64, emailAddr string) (*entity_email.Email, error) {
	ee := entity_email.New(i.emailRepository)
	email, err := ee.Update(emailID, emailAddr)
	if err != nil {
		return nil, err
	}
	return email, nil
}
