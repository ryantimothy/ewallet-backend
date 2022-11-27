package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserId   int    `gorm:"primaryKey;column:user_id"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`

	WalletId   int    `gorm:"column:wallet_id"`
	UserWallet Wallet `gorm:"foreignkey:WalletId;references:WalletId"`

	CreatedAt time.Time      `gorm:"column:created_at`
	UpdatedAt time.Time      `gorm:"column:updated_at`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at`
}

type UserDetailResponse struct {
	UserID     int
	Email      string
	UserWallet Wallet
}

type UserRegisterResponse struct {
	UserID   int
	Email    string
	WalletID int
}
