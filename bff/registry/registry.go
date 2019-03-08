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
	UserUseCase  usecase_user.UseCase
	EmailUseCase usecase_email.UseCase
	Context      context.Context
	EmailService microservice_email.Service
}

func New(db *gorm.DB, ctx context.Context, emailConn *grpc.ClientConn) *Registry {
	ur := repository_user.New(db)
	pr := repository_profile.New(db)
	er := repository_email.New(db)

	return &Registry{
		UserUseCase:  interactor_user.New(ur, pr, er),
		EmailUseCase: interactor_email.New(er),
		Context:      ctx,
		EmailService: microservice_email.New(emailConn),
	}
}
