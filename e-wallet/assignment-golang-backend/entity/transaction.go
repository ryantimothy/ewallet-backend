package entity

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	TransactionId   int    `gorm:"primaryKey;column:transaction_id"`
	TransactionType string `gorm:"column:transaction_type" `
	SenderID        int    `gorm:"column:sender_id" `
	ReceiverID      int    `gorm:"column:receiver_id" `
	Amount          int    `gorm:"column:amount" `
	Description     string `gorm:"column:description" `
	SourceOfFundID  int    `gorm:"column:source_fund_id" `

	CreatedAt time.Time      `gorm:"column:created_at`
	UpdatedAt time.Time      `gorm:"column:updated_at`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at`
}

type TopUpRequest struct {
	Amount         int
	SourceOfFundID int
}

type TransferRequest struct {
	Recipient   int
	Amount      int
	Description string
}

type Query struct {
	SortBy string
	Sort   string
	Desc   string
}
