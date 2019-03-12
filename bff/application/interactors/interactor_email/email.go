package interactor_email

import (
	"context"
	"go-bff/bff/domain/entities/entity_email"
	"go-bff/bff/domain/entities/entity_notify"
	"go-bff/notify/pb"
)

type interactor struct {
	emailRepository entity_email.Repository
	notifyService   entity_notify.Service
	context         context.Context
}

func New(er entity_email.Repository, ns entity_notify.Service, ctx context.Context) *interactor {
	return &interactor{emailRepository: er, notifyService: ns, context: ctx}
}

func (i *interactor) First(emailID uint64) (*entity_email.Email, error) {
	ee := entity_email.New(i.emailRepository)
	email, err := ee.First(emailID)
	if err != nil {
		return nil, err
	}

	res, err := i.notifyService.Reverse(i.context, &pb.ReverseRequest{Email: email.Email})
	if err != nil {
		return nil, err
	}
	email.Email = res.EmailAddress

	return email, nil
}

func (i *interactor) Update(emailID uint64, emailAddr string) (*entity_email.Email, error) {
	res, err := i.notifyService.Build(i.context, &pb.BuildRequest{Email: emailAddr})
	if err != nil {
		return nil, err
	}

	ee := entity_email.New(i.emailRepository)
	email, err := ee.Update(emailID, res.EmailAddress)
	if err != nil {
		return nil, err
	}
	return email, nil
}
