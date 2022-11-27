package usecase

import (
	"errors"
	"ewallet/entity"
	"ewallet/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionUsecase_GetAll_Success(t *testing.T) {
	tr := mocks.NewTransactionRepo(t)
	wr := mocks.NewWalletRepo(t)
	ur := mocks.NewUserRepo(t)

	u := NewTransactionUsecase(tr, wr, ur)

	entityTrans := entity.Transaction{
		TransactionId:   1,
		TransactionType: "TOP UP",
		SenderID:        1,
		ReceiverID:      2,
		Amount:          10000,
		Description:     "Coffee",
		SourceOfFundID:  1,
	}
	entityQ := entity.Query{
		SortBy: "test",
		Sort:   "test",
		Desc:   "test",
	}
	tr.On("GetAll", 1, entityQ).
		Return([]*entity.Transaction{&entityTrans}, nil)

	r, err := u.GetAll(1, entityQ)

	assert.Nil(t, err)
	assert.NotNil(t, r)
}

func TestTransactionUsecase_GetAll_Fail(t *testing.T) {
	tr := mocks.NewTransactionRepo(t)
	wr := mocks.NewWalletRepo(t)
	ur := mocks.NewUserRepo(t)

	u := NewTransactionUsecase(tr, wr, ur)

	entityQ := entity.Query{
		SortBy: "test",
		Sort:   "test",
		Desc:   "test",
	}
	tr.On("GetAll", 1, entityQ).
		Return(nil, errors.New("Error: invalid input"))

	r, err := u.GetAll(1, entityQ)

	assert.NotNil(t, err)
	assert.Nil(t, r)
}

func TestTopUpDescription(t *testing.T) {
	type args struct {
		SoF int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "test 1",
			args: args{
				SoF: 1,
			},
			want:    "Top Up from Bank Transfer",
			wantErr: nil,
		},
		{
			name: "test 2",
			args: args{
				SoF: 2,
			},
			want:    "Top Up from Credit Card",
			wantErr: nil,
		},
		{
			name: "test 3",
			args: args{
				SoF: 3,
			},
			want:    "Top Up from Cash",
			wantErr: nil,
		},
		{
			name: "test 4",
			args: args{
				SoF: 4,
			},
			want:    "",
			wantErr: errors.New("invalid source of fund"),
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := TopUpDescription(tt.args.SoF)

			if i < 3 {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
