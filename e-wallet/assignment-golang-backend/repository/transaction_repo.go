package repository

import (
	"errors"
	"ewallet/entity"
	"fmt"

	"gorm.io/gorm"
)

type TransactionRepo interface {
	GetAll(int, entity.Query) ([]*entity.Transaction, error)
	TopUp(*entity.Transaction) error
	Transfer(tr *entity.Transaction) error
}

type transactionImpl struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) TransactionRepo {
	return &transactionImpl{
		db: db,
	}

}

func (t *transactionImpl) GetAll(id int, q entity.Query) ([]*entity.Transaction, error) {
	var tr []*entity.Transaction

	orderString := q.SortBy + " " + q.Sort
	fmt.Println(orderString)

	if err := t.db.Limit(10).Order(orderString).Where("(sender_id = ? OR receiver_id = ?) AND description ILIKE ?", id, id, q.Desc).Find(&tr).Error; err != nil {

		return nil, errors.New("invalid input")
	}

	return tr, nil
}

func (t *transactionImpl) TopUp(tr *entity.Transaction) error {

	if err := t.db.Create(&tr).Error; err != nil {

		return err
	}

	return nil
}

func (t *transactionImpl) Transfer(tr *entity.Transaction) error {

	if err := t.db.Create(&tr).Error; err != nil {

		return err
	}

	return nil
}
