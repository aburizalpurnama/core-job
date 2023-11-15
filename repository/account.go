package repository

import (
	"core-users-job/model"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type (
	AccountRepository interface {
		SelectOpenAccounts(date time.Time) (reports []model.OpenAccountReport, err error)
		SelectCloseAccounts(date time.Time) (reports []model.CloseAccountReport, err error)
	}

	accountRepositoryImpl struct {
		db *gorm.DB
	}
)

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepositoryImpl{
		db: db,
	}
}

var (
	selectOpenAccounts = `SELECT cif, account_number, sanction_limit, expired_date, bank_account_number, scheme_code, due_date, cycle_difference, reference_id, status_date AS open_account_date
	FROM users.accounts a 
	WHERE a.deleted = 0 AND a.status = 'active' AND (a.status_date BETWEEN ? AND ?)
	ORDER BY id ASC;`
	selectCloseAccounts = `SELECT cif, account_number, sanction_limit, block_code, bank_account_number, status_date AS close_account_date
	FROM users.accounts a 
	WHERE a.deleted = 0 AND a.status = 'closed' AND (a.status_date BETWEEN ? AND ?)
	ORDER BY id ASC;`
)

func (r *accountRepositoryImpl) SelectOpenAccounts(date time.Time) (reports []model.OpenAccountReport, err error) {
	year, m, day := date.Date()
	start := time.Date(year, m, day, 0, 0, 0, 0, date.Location())
	end := time.Date(year, m, day, 23, 59, 59, 59, date.Location())

	db := r.db.Raw(selectOpenAccounts, start, end)
	if err = db.Scan(&reports).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Errorf("REPO ACCOUNT -> SelectOpenAccounts, err: %s", err)
	}

	return
}

func (r *accountRepositoryImpl) SelectCloseAccounts(date time.Time) (reports []model.CloseAccountReport, err error) {
	year, m, day := date.Date()
	start := time.Date(year, m, day, 0, 0, 0, 0, date.Location())
	end := time.Date(year, m, day, 23, 59, 59, 59, date.Location())

	db := r.db.Raw(selectCloseAccounts, start, end)
	if err = db.Scan(&reports).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Errorf("REPO ACCOUNT -> SelectCloseAccounts, err: %s", err)
	}

	return
}
