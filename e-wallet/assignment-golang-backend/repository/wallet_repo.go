package repository

import (
	"errors"
	"ewallet/entity"
	"time"

	"gorm.io/gorm"
)

type WalletRepo interface {
	CreateWallet() (*entity.Wallet, error)
	AddBalance(walletID int, amount int) error
	ReduceBalance(walletID int, amount int) error
	GetUserByWalletID(walletid int) (*entity.User, error)
	GetWallet(walletID int) (*entity.Wallet, error)
}

type WalletImpl struct {
	db *gorm.DB
}

func NewWalletRepo(db *gorm.DB) WalletRepo {
	return &WalletImpl{
		db: db,
	}

}

func (w *WalletImpl) CreateWallet() (*entity.Wallet, error) {
	wallet := &entity.Wallet{
		Balance:   0,
		CreatedAt: time.Now(),
	}

	if err := w.db.Create(&wallet).Error; err != nil {
		return nil, err
	}

	return wallet, nil

}

func (u *WalletImpl) GetUserByWalletID(walletid int) (*entity.User, error) {
	user := entity.User{}

	if err := u.db.Where("wallet_id = ?", walletid).First(&user).Error; err != nil {

		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (w *WalletImpl) GetWallet(walletID int) (*entity.Wallet, error) {
	wallet := entity.Wallet{}

	if err := w.db.First(&wallet, walletID).Error; err != nil {

		return &entity.Wallet{}, errors.New("wallet not found")
	}

	return &wallet, nil
}

func (w *WalletImpl) AddBalance(walletID int, amount int) error {
	wallet := entity.Wallet{}

	if err := w.db.First(&wallet, walletID).Error; err != nil {

		return errors.New("wallet not found")
	}

	w.db.Model(&wallet).UpdateColumn("balance", gorm.Expr("balance + ?", amount))

	return nil
}

func (w *WalletImpl) ReduceBalance(walletID int, amount int) error {
	wallet := entity.Wallet{}

	if err := w.db.First(&wallet, walletID).Error; err != nil {

		return errors.New("wallet not found")
	}

	w.db.Model(&wallet).UpdateColumn("balance", gorm.Expr("balance - ?", amount))

	return nil
}
