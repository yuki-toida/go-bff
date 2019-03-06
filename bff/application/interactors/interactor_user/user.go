package interactor_user

import (
	"go-bff/bff/domain/entities/entity_email"
	"go-bff/bff/domain/entities/entity_profile"
	"go-bff/bff/domain/entities/entity_user"
)

type interactor struct {
	userRepository    entity_user.Repository
	profileRepository entity_profile.Repository
	emailRepository   entity_email.Repository
}

func New(ur entity_user.Repository, pr entity_profile.Repository, er entity_email.Repository) *interactor {
	return &interactor{userRepository: ur, profileRepository: pr, emailRepository: er}
}

func (i *interactor) Find() ([]entity_user.User, error) {
	user := entity_user.New(i.userRepository)
	return user.Find()
}

func (i *interactor) First(id uint64) (*entity_user.User, error) {
	user := entity_user.New(i.userRepository)
	return user.First(id)
}

func (i *interactor) Create(name string) (*entity_user.User, error) {
	profile := entity_profile.New(i.profileRepository)
	p, err := profile.Create(name)
	if err != nil {
		return nil, err
	}

	user := entity_user.New(i.userRepository)
	return user.Create(p)
}

func (i *interactor) Delete(userID uint64) error {
	user := entity_user.New(i.userRepository)
	return user.Delete(userID)
}

func (i *interactor) CreateEmail(userID uint64, emailAddr string) (*entity_user.User, error) {
	eu := entity_user.New(i.userRepository)
	user, err := eu.First(userID)
	if err != nil {
		return nil, err
	}
	ee := entity_email.New(i.emailRepository)
	email, err := ee.Create(userID, emailAddr)
	if err != nil {
		return nil, err
	}
	eu.AddEmail(user, email)
	return user, nil
}
