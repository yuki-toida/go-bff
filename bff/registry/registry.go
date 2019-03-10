package registry

import (
	"context"
	"github.com/jinzhu/gorm"
	"go-bff/bff/adapter/microservices/microservice_email"
	"go-bff/bff/adapter/repositories/repository_email"
	"go-bff/bff/adapter/repositories/repository_profile"
	"go-bff/bff/adapter/repositories/repository_user"
	"go-bff/bff/application/interactors/interactor_email"
	"go-bff/bff/application/interactors/interactor_user"
	"go-bff/bff/application/usecase/usecase_email"
	"go-bff/bff/application/usecase/usecase_user"
	"google.golang.org/grpc"
)

type Registry struct {
	Context       context.Context
	UseCases      useCases
	MicroServices microServices
}

type useCases struct {
	User  usecase_user.UseCase
	Email usecase_email.UseCase
}

type microServices struct {
	Email microservice_email.Service
}

func New(db *gorm.DB, ctx context.Context, emailConn *grpc.ClientConn) *Registry {
	ur := repository_user.New(db)
	pr := repository_profile.New(db)
	er := repository_email.New(db)

	return &Registry{
		Context: ctx,
		UseCases: useCases{
			User:  interactor_user.New(ur, pr, er),
			Email: interactor_email.New(er),
		},
		MicroServices: microServices{
			Email: microservice_email.New(emailConn),
		},
	}
}
