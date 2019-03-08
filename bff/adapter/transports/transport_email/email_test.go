package transport_email

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go-bff/bff/adapter/microservices/microservice_email/mock_microservice_email"
	"go-bff/bff/application/usecase/usecase_email/mock_usecase_email"
	"go-bff/bff/domain/entities"
	"go-bff/bff/domain/entities/entity_email"
	"go-bff/email/pb"
	"testing"
)

func TestMakeFirstEndpoint(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	email := &entity_email.Email{
		entities.Model{ID: 1},
		"hoge",
		1,
	}
	mu := mock_usecase_email.NewMockUseCase(ctrl)
	mu.EXPECT().First(uint64(1)).Return(email, nil)

	reversed := "egoh"
	ctx := context.Background()
	ms := mock_microservice_email.NewMockService(ctrl)
	ms.EXPECT().Reverse(ctx, &pb.ReverseRequest{Email: email.Email}).Return(&pb.ReverseResponse{EmailAddress: reversed}, nil)

	ep := MakeFirstEndpoint(mu, ms, ctx)
	res, err := ep(ctx, "1")
	email.Email = reversed
	assert.Nil(t, err)
	assert.Equal(t, res, email)
}
