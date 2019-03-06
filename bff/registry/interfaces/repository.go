package interfaces

import (
	"go-bff/bff/domain/entities/entity_email"
	"go-bff/bff/domain/entities/entity_profile"
	"go-bff/bff/domain/entities/entity_user"
)

type Repositories interface {
	NewUserRepository() entity_user.Repository
	NewProfileRepository() entity_profile.Repository
	NewEmailRepository() entity_email.Repository
}
