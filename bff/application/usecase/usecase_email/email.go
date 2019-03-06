package usecase_email

import (
	"go-bff/bff/domain/entities/entity_email"
)

type UseCase interface {
	First(emailID uint64) (*entity_email.Email, error)
	Update(emailID uint64, emailAddr string) (*entity_email.Email, error)
}
