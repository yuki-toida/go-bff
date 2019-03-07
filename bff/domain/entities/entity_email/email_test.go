package entity_email_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go-bff/bff/domain/entities"
	"go-bff/bff/domain/entities/entity_email"
	"go-bff/bff/domain/entities/entity_email/mock_email"
	"testing"
)

// mock作成コマンド
// mkdir mock_email
// mockgen go-bff/bff/domain/entities/entity_email Repository > mock_email/mock_email_repository.go
func TestFirst(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_entity_email.NewMockRepository(ctrl)
	m.EXPECT().First(uint64(1)).Return(&entity_email.Email{
		entities.Model{ID: 1},
		"1@hacobu.jp",
		1,
	}, nil)

	ee := entity_email.New(m)
	email, err := ee.First(uint64(1))
	assert.Nil(t, err)
	assert.Equal(t, &entity_email.Email{
		entities.Model{ID: 1},
		"1@hacobu.jp",
		1,
	}, email)
}
