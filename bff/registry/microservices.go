package registry

import (
	"go-bff/bff/adapter/microservices/microservice_email"
	"go-bff/bff/registry/interfaces"
	"google.golang.org/grpc"
)

type microservices struct {
	emailConn *grpc.ClientConn
}

func NewMicroServices(emailConn *grpc.ClientConn) interfaces.MicroServices {
	return &microservices{emailConn: emailConn}
}

func (m *microservices) GetEmailService() microservice_email.Service {
	return microservice_email.New(m.emailConn)
}
