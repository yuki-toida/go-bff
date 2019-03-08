#!/bin/sh

# MicroServices
mockgen go-bff/bff/adapter/microservices/microservice_email Service > bff/adapter/microservices/microservice_email/mock_microservice_email/mock_email.go

# UseCase
mockgen go-bff/bff/application/usecase/usecase_email UseCase > bff/application/usecase/usecase_email/mock_usecase_email/mock_email.go
