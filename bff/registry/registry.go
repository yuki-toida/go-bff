package registry

import (
	"context"
	"github.com/jinzhu/gorm"
	"go-bff/bff/adapter/gateways/api/api_notify"
	"go-bff/bff/adapter/gateways/db/db_email"
	"go-bff/bff/adapter/gateways/db/db_profile"
	"go-bff/bff/adapter/gateways/db/db_user"
	"go-bff/bff/application/interactors/interactor_email"
	"go-bff/bff/application/interactors/interactor_user"
	"go-bff/bff/application/usecase/usecase_email"
	"go-bff/bff/application/usecase/usecase_user"
	"google.golang.org/grpc"
)

type Registry struct {
	UserUseCase  usecase_user.UseCase
	EmailUseCase usecase_email.UseCase
}

func New(db *gorm.DB, ctx context.Context, notifyConn *grpc.ClientConn) *Registry {
	ur := db_user.New(db)
	pr := db_profile.New(db)
	er := db_email.New(db)

	ns := api_notify.New(notifyConn)

	return &Registry{
		UserUseCase:  interactor_user.New(ur, pr, er),
		EmailUseCase: interactor_email.New(er, ns, ctx),
	}
}
