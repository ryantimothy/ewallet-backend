package usecase

import (
	"errors"
	"ewallet/entity"
	"ewallet/repository"
)

const (
	minTopUp      = 50000
	maxTopUp      = 10000000
	minTransfer   = 1000
	maxTransfer   = 50000000
	maxLenDesc    = 35
	unavailableID = 0
)

type TransactionUsecase interface {
	GetAll(int, entity.Query) ([]*entity.Transaction, error)
	TopUp(req entity.TopUpRequest, id int) error
	Transfer(req entity.TransferRequest, id int) error
}

type transactionUsecaseImpl struct {
	repository       repository.TransactionRepo
	walletrepository repository.WalletRepo
	userrepository   repository.UserRepo
}

func NewTransactionUsecase(r repository.TransactionRepo, wr repository.WalletRepo, ur repository.UserRepo) TransactionUsecase {
	return &transactionUsecaseImpl{
		repository:       r,
		walletrepository: wr,
		userrepository:   ur,
	}
}

func (tu *transactionUsecaseImpl) GetAll(id int, q entity.Query) ([]*entity.Transaction, error) {
	resp, err := tu.repository.GetAll(id, q)
	if err != nil {

		return nil, err
	}

	return resp, nil
}

func TopUpDescription(SoF int) (string, error) {
	switch SoF {
	case 1:

		return "Top Up from Bank Transfer", nil

	case 2:

		return "Top Up from Credit Card", nil

	case 3:

		return "Top Up from Cash", nil

	default:

		return "", errors.New("invalid source of fund")
	}
}

func (tu *transactionUsecaseImpl) TopUp(req entity.TopUpRequest, id int) error {
	if req.Amount < minTopUp || req.Amount > maxTopUp {

		return errors.New("invalid topup amount")
	}

	desc, err := TopUpDescription(req.SourceOfFundID)
	if err != nil {

		return err
	}

	user, err := tu.userrepository.GetUser(id)
	if err != nil {

		return err

	}

	err = tu.walletrepository.AddBalance(user.WalletId, req.Amount)
	if err != nil {

		return err
	}

	transaction := &entity.Transaction{}
	transaction.TransactionType = "TOP UP"
	transaction.SenderID = id
	transaction.ReceiverID = id
	transaction.Amount = req.Amount
	transaction.Description = desc
	transaction.SourceOfFundID = req.SourceOfFundID

	err = tu.repository.TopUp(transaction)
	if err != nil {

		return err
	}

	return nil
}

func (tu *transactionUsecaseImpl) Transfer(req entity.TransferRequest, id int) error {
	if req.Amount < minTransfer || req.Amount > maxTransfer {

		return errors.New("invalid transfer amount")
	}

	if len(req.Description) > maxLenDesc {

		return errors.New("description exceeds maximum character length")
	}

	userSender, err := tu.userrepository.GetUser(id)
	if err != nil {

		return err

	}

	userReceiver, err := tu.walletrepository.GetUserByWalletID(req.Recipient)
	if err != nil {

		return err

	}

	senderWallet, err := tu.walletrepository.GetWallet(userSender.WalletId)
	if err != nil {

		return err

	}
	if (senderWallet.Balance - req.Amount) < 0 {

		return errors.New("insufficient balance")
	}

	err = tu.walletrepository.ReduceBalance(userSender.WalletId, req.Amount)
	if err != nil {

		return err

	}

	err = tu.walletrepository.AddBalance(req.Recipient, req.Amount)
	if err != nil {

		return err

	}

	transaction := &entity.Transaction{}
	transaction.TransactionType = "TRANSFER"
	transaction.SenderID = id
	transaction.ReceiverID = userReceiver.UserId
	transaction.Amount = req.Amount
	transaction.Description = req.Description
	transaction.SourceOfFundID = unavailableID

	err = tu.repository.Transfer(transaction)
	if err != nil {

		return err
	}

	return nil

}
