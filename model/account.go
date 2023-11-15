package model

import (
	database "core-users-service/databases/rdbms"
	"time"
)

const TableNameAccount = "users.accounts"

// Account mapped from table <accounts>
type Account struct {
	database.BaseModel
	CifID              int32       `gorm:"column:cif_id;not null" json:"cif_id"`
	Cif                int64       `gorm:"column:cif;not null" json:"cif"`
	CifName            string      `gorm:"column:cif_name;not null" json:"cif_name"`
	AccountNumber      int64       `gorm:"column:account_number;not null" json:"account_number"`
	SanctionLimit      float64     `gorm:"column:sanction_limit;not null" json:"sanction_limit"`
	AvailableLimit     float64     `gorm:"column:available_limit;not null" json:"available_limit"`
	HashLimit          string      `gorm:"column:hash_limit" json:"hash_limit"`
	ExpiredDate        time.Time   `gorm:"column:expired_date;not null" json:"expired_date"`
	BankAccountNumber  string      `gorm:"column:bank_account_number;not null" json:"bank_account_number"`
	SchemeCode         string      `gorm:"column:scheme_code;not null" json:"scheme_code"`
	DueDate            int32       `gorm:"column:due_date;not null" json:"due_date"`
	CycleDifference    int32       `gorm:"column:cycle_difference;not null" json:"cycle_difference"`
	BillDate           int32       `gorm:"column:bill_date;not null" json:"bill_date"`
	AccountClose       int32       `gorm:"column:account_close;not null;comment:status account" json:"account_close"` // status account
	ReferenceID        string      `gorm:"column:reference_id;not null" json:"reference_id"`
	Collectibility     int32       `gorm:"column:collectibility;not null;default:1" json:"collectibility"`
	BlockCode          *string     `gorm:"column:block_code" json:"block_code"`
	TemporaryBlockCode string      `gorm:"column:temporary_block_code" json:"temporary_block_code"`
	Status             string      `gorm:"column:status" json:"status"`
	StatusDate         time.Time   `gorm:"column:status_date" json:"status_date"`
	CreatedBy          interface{} `gorm:"column:created_by;type:json;not null;comment:{from: api/dashboard, admin_id, admin_name}" json:"created_by"` // {from: api/dashboard, admin_id, admin_name}
}

type OpenAccountReport struct {
	AccountNumber     string    `gorm:"column:account_number" json:"account_number"`
	Cif               int64     `gorm:"column:cif" json:"cif"`
	SanctionLimit     float64   `gorm:"column:sanction_limit" json:"sanction_limit"`
	ExpiredDate       time.Time `gorm:"column:expired_date" json:"expired_date"`
	BankAccountNumber string    `gorm:"column:bank_account_number" json:"bank_account_number"`
	SchemeCode        string    `gorm:"column:scheme_code" json:"scheme_code"`
	DueDate           int64     `gorm:"column:due_date" json:"due_date"`
	CycleDifference   int32     `gorm:"column:cycle_difference" json:"cycle_difference"`
	ReferenceID       string    `gorm:"column:reference_id" json:"reference_id"`
	OpenAccountDate   time.Time `gorm:"column:open_account_date" json:"open_account_date"`
}

type CloseAccountReport struct {
	AccountNumber     string    `gorm:"column:account_number" json:"account_number"`
	Cif               int64     `gorm:"column:cif" json:"cif"`
	SanctionLimit     float64   `gorm:"column:sanction_limit" json:"sanction_limit"`
	BlockCode         string    `gorm:"column:block_code" json:"block_code"`
	BankAccountNumber string    `gorm:"column:bank_account_number" json:"bank_account_number"`
	CloseAccountDate  time.Time `gorm:"column:close_account_date" json:"close_account_date"`
}

// TableName Account's table name
func (*Account) TableName() string {
	return TableNameAccount
}
