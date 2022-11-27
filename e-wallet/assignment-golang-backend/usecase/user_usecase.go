package usecase

import (
	"errors"
	"ewallet/entity"
	"ewallet/repository"
	"ewallet/utils"
)

type UserUsecase interface {
	Register(*entity.User) (*entity.UserRegisterResponse, error)
	Login(email, password string) (string, error)
	GetDetail(id int) (*entity.UserDetailResponse, error)
}

type userUsecaseImpl struct {
	repository       repository.UserRepo
	walletrepository repository.WalletRepo
}

func NewUserUsecase(repository repository.UserRepo, walletrepository repository.WalletRepo) UserUsecase {
	return &userUsecaseImpl{
		repository:       repository,
		walletrepository: walletrepository,
	}
}

func CreateUserRegisterResponse(resp *entity.User) *entity.UserRegisterResponse {
	response := entity.UserRegisterResponse{}

	response.Email = resp.Email
	response.UserID = resp.UserId
	response.WalletID = resp.WalletId

	return &response
}

func (u *userUsecaseImpl) Register(user *entity.User) (*entity.UserRegisterResponse, error) {
	temp, err := utils.HashAndSalt(user.Password)

	if user.Email == "" {

		return nil, errors.New("empty email")
	}

	if !utils.ValidateEmail(user.Email) {

		return nil, errors.New("invalid email")
	}

	if user.Password == "" {

		return nil, errors.New("empty password")
	}

	if err != nil {

		return nil, err
	}

	user.Password = temp

	wallet, err := u.walletrepository.CreateWallet()
	if err != nil {

		return nil, err
	}

	user.WalletId = wallet.WalletId
	r, err := u.repository.Register(user)
	if err != nil {

		return nil, err
	}

	resp := CreateUserRegisterResponse(r)

	return resp, nil
}

func (u *userUsecaseImpl) Login(email, password string) (string, error) {
	user, err := u.repository.GetUsername(email)
	if err != nil {

		return "", err
	}

	if !utils.ComparePassword(user.Password, password) {
		return "", errors.New("wrong username or password")
	}

	token, err := utils.GenerateToken(uint(user.UserId))
	if err != nil {

		return "", err
	}

	return token, nil

}

func (u *userUsecaseImpl) GetDetail(id int) (*entity.UserDetailResponse, error) {
	resp, err := u.repository.GetDetail(id)

	if err != nil {

		return nil, err
	}

	return resp, nil
}
