package main

import (
	"core-users-job/handler"
	"core-users-job/repository"
	"core-users-job/router"
	"core-users-job/usecase"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func main() {
	loadEnv()
	run()
}

func run() {
	db := connectDB()
	accountRepo := repository.NewAccountRepository(db)
	accountUsecase := usecase.NewAccountUsecase(accountRepo)
	reportHandler := handler.NewReportHandler(accountUsecase)
	app := router.InitRoutes(fiber.Config{AppName: os.Getenv("APP_NAME")}, reportHandler)
	if err := app.Listen(fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))); err != nil {
		log.Fatal(err)
	}
}

func loadEnv() {
	// TODO: combine with https://github.com/caarlos0/env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}

func connectDB() *gorm.DB {
	dsnMaster := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s", os.Getenv("DB_MASTER_HOST"), os.Getenv("DB_MASTER_PORT"), os.Getenv("DB_MASTER_USER"), os.Getenv("DB_MASTER_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_SSL_MODE"), os.Getenv("DB_TIME_ZONE"))
	dsnSlave := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s", os.Getenv("DB_SLAVE_HOST"), os.Getenv("DB_SLAVE_PORT"), os.Getenv("DB_SLAVE_USER"), os.Getenv("DB_SLAVE_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_SSL_MODE"), os.Getenv("DB_TIME_ZONE"))
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
