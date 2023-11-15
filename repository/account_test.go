package repository

import (
	"fmt"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func TestSelectOpenAccounts(t *testing.T) {
	db := connectDB()
	repo := NewAccountRepository(db)
	openReports, err := repo.SelectOpenAccounts(time.Now())
	assert.NoError(t, err)

	closeReports, err := repo.SelectCloseAccounts(time.Now())
	assert.NoError(t, err)

	fmt.Printf("openReports: %v\n", openReports)
	fmt.Printf("closeReports: %v\n", closeReports)
	fmt.Printf("err: %v\n", err)
}

func connectDB() *gorm.DB {
	dsnMaster := "host=localhost port=5432 user=admin password=secret dbname=briceria sslmode=disable TimeZone=Asia/Jakarta"
	dsnSlave := "host=localhost port=5432 user=admin password=secret dbname=briceria sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsnMaster), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{postgres.Open(dsnMaster)},
		Replicas: []gorm.Dialector{postgres.Open(dsnSlave)},
		Policy:   dbresolver.RandomPolicy{},
	}))
	if err != nil {
		log.Fatalf("[DB SRV] Error Connection Testing to DB - %v", err)
	}

	log.Info("[DB SRV] Successful Connection Testing to DB")

	return db
}
