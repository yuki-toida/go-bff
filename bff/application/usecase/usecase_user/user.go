package usecase_user

import (
	"go-bff/bff/domain/entities/entity_user"
)

type UseCase interface {
	Find() ([]entity_user.User, error)
	First(userID uint64) (*entity_user.User, error)
	Create(name string) (*entity_user.User, error)
	Delete(userID uint64) error
	CreateEmail(userID uint64, emailAddr string) (*entity_user.User, error)
}
