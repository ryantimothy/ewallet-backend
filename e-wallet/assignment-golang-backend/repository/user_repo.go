package repository

import (
	"errors"
	"ewallet/entity"

	"gorm.io/gorm"
)

type UserRepo interface {
	Register(*entity.User) (*entity.User, error)
	GetUsername(email string) (*entity.User, error)
	GetUser(id int) (*entity.User, error)
	GetDetail(id int) (*entity.UserDetailResponse, error)
}

type UserImpl struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &UserImpl{
		db: db,
	}

}
func (u *UserImpl) Register(user *entity.User) (*entity.User, error) {

	if err := u.db.Where("email LIKE ?", user.Email).First(&user).Error; err == nil {

		return &entity.User{}, errors.New("user has existed")
	}

	if err := u.db.Create(&user).Error; err != nil {

		return &entity.User{}, err
	}

	return user, nil
}
func (u *UserImpl) GetUsername(email string) (*entity.User, error) {
	user := entity.User{}

	if err := u.db.Where("email ILIKE ?", email).First(&user).Error; err != nil {

		return nil, errors.New("user not found")
	}

	return &user, nil

}

func (u *UserImpl) GetUser(id int) (*entity.User, error) {
	user := entity.User{}

	if err := u.db.Where("user_id = ?", id).First(&user).Error; err != nil {

		return nil, errors.New("user not found")
	}

	return &user, nil
}

func CreateUserDetailResponse(user entity.User) entity.UserDetailResponse {
	resp := entity.UserDetailResponse{
		UserID:     user.UserId,
		Email:      user.Email,
		UserWallet: user.UserWallet,
	}

	return resp
}

func (u *UserImpl) GetDetail(id int) (*entity.UserDetailResponse, error) {
	var user *entity.User

	if err := u.db.Where("user_id = ?", id).Preload("UserWallet").Find(&user).Error; err != nil {

		return nil, err
	}

	resp := CreateUserDetailResponse(*user)

	return &resp, nil

}
