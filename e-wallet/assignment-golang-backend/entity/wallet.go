package entity

import (
	"time"

	"gorm.io/gorm"
)

type Wallet struct {
	WalletId  int            `gorm:"primaryKey;column:wallet_id"`
	Balance   int            `gorm:"column:balance"`
	CreatedAt time.Time      `gorm:"column:created_at`
	UpdatedAt time.Time      `gorm:"column:updated_at`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at`
}
