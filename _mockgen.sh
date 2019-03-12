#!/bin/sh

# MicroServices
mockgen go-bff/bff/domain/entities/entity_notify Service > bff/domain/entities/entity_notify/mock_entity_notify/mock_notify_service.go

# UseCase
mockgen go-bff/bff/application/usecase/usecase_email UseCase > bff/application/usecase/usecase_email/mock_usecase_email/mock_email.go
