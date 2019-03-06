package registry

import (
	"github.com/jinzhu/gorm"
	"go-bff/bff/adapter/repositories/repository_email"
	"go-bff/bff/adapter/repositories/repository_profile"
	"go-bff/bff/adapter/repositories/repository_user"
	"go-bff/bff/domain/entities/entity_email"
	"go-bff/bff/domain/entities/entity_profile"
	"go-bff/bff/domain/entities/entity_user"
	"go-bff/bff/registry/interfaces"
)

type repositories struct {
	db *gorm.DB
}

func NewRepositories(db *gorm.DB) interfaces.Repositories {
	return &repositories{db: db}
}

func (r *repositories) NewUserRepository() entity_user.Repository {
	return repository_user.New(r.db)
}

func (r *repositories) NewProfileRepository() entity_profile.Repository {
	return repository_profile.New(r.db)
}

func (r *repositories) NewEmailRepository() entity_email.Repository {
	return repository_email.New(r.db)
}
