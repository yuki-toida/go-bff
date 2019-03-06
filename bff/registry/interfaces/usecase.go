package interfaces

import (
	"go-bff/bff/application/usecase/usecase_email"
	"go-bff/bff/application/usecase/usecase_user"
)

type UseCases interface {
	NewUserUseCase() usecase_user.UseCase
	NewEmailUseCase() usecase_email.UseCase
}
