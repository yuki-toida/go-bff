package interfaces

import "go-bff/bff/adapter/microservices/microservice_email"

type MicroServices interface {
	GetEmailService() microservice_email.Service
}
