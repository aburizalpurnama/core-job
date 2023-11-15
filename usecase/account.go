package usecase

import (
	"core-users-job/repository"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

type (
	AccountUsecase interface {
		GenerateOpenAccountReport() (err error)
	}

	accountUsecaseImpl struct {
		accountRepo repository.AccountRepository
	}
)

func NewAccountUsecase(accountRepo repository.AccountRepository) AccountUsecase {
	return &accountUsecaseImpl{
		accountRepo: accountRepo,
	}
}

func (u *accountUsecaseImpl) GenerateOpenAccountReport() (err error) {
	date := time.Now().Local().AddDate(0, 0, -1)

	accounts, err := u.accountRepo.SelectOpenAccounts(date)
	if err != nil {
		return err
	}

	year, m, day := date.Date()

	rootDir := "/open-account-report"
	yearDir := fmt.Sprintf("%d", year)
	monthDir := m.String()
	// path := fmt.Sprintf("%s%s/%s/%s/%d-%d-%d.csv", os.Getenv("REPORT_DIR"), rootDir, yearDir, monthDir, year, m, day)

	// file, err := os.Create(path)

	path := fmt.Sprintf("%s/%s/%s/%d-%d-%d.csv", rootDir, yearDir, monthDir, year, m, day)
	file, err := create(path)
	if err != nil {
		log.Errorf("failed to open file, err:%v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var data [][]string

	data = append(data, []string{"account_number", "cif", "sanction_limit", "expired_date", "bank_account_number", "scheme_code", "due_date", "cycle_difference", "reference_id"})
	for _, a := range accounts {
		row := []string{a.AccountNumber, fmt.Sprintf("%d", a.Cif), fmt.Sprintf("%f", a.SanctionLimit), a.ExpiredDate.Local().Format("2006-01-02"), a.BankAccountNumber, a.SchemeCode, fmt.Sprintf("%d", a.DueDate), fmt.Sprintf("%d", a.CycleDifference), a.ReferenceID}
		data = append(data, row)
	}

	err = writer.WriteAll(data)
	if err != nil {
		log.Errorf("failed to write csv, err:%v", err)
	}

	return
}

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}
