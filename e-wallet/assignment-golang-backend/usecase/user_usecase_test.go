package usecase

import (
	"errors"
	"ewallet/entity"
	"ewallet/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserUsecase_GetDetail_Success(t *testing.T) {

	wr := mocks.NewWalletRepo(t)
	ur := mocks.NewUserRepo(t)

	u := NewUserUsecase(ur, wr)

	resp := &entity.UserDetailResponse{
		UserID: 1,
		Email:  "test@gmail.com",
	}

	ur.On("GetDetail", 1).
		Return(resp, nil)

	r, err := u.GetDetail(1)

	assert.Nil(t, err)
	assert.NotNil(t, r)
}

func TestUserUsecase_GetDetail_Error(t *testing.T) {

	wr := mocks.NewWalletRepo(t)
	ur := mocks.NewUserRepo(t)

	u := NewUserUsecase(ur, wr)

	ur.On("GetDetail", 1).
		Return(nil, errors.New(""))

	r, err := u.GetDetail(1)

	assert.NotNil(t, err)
	assert.Nil(t, r)
}
