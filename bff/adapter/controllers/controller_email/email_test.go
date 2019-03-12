package controller_email

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go-bff/bff/application/usecase/usecase_email/mock_usecase_email"
	"go-bff/bff/domain/entities"
	"go-bff/bff/domain/entities/entity_email"
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

	ctx := context.Background()

	ep := MakeFirstEndpoint(mu)
	res, err := ep(ctx, "1")
	assert.Nil(t, err)
	assert.Equal(t, res, email)
}
