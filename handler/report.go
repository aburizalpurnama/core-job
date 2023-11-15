package handler

import (
	"core-users-job/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type (
	ReportHandler interface {
		GenerateOpenAccountReport(c *fiber.Ctx) error
	}

	reportHandlerImpl struct {
		accountUsecase usecase.AccountUsecase
	}
)

func NewReportHandler(accountUsecase usecase.AccountUsecase) ReportHandler {
	return &reportHandlerImpl{
		accountUsecase: accountUsecase,
	}
}

func (h *reportHandlerImpl) GenerateOpenAccountReport(c *fiber.Ctx) error {
	err := h.accountUsecase.GenerateOpenAccountReport()
	switch err {
	case nil:
		return c.SendStatus(fiber.StatusCreated)

	case gorm.ErrRecordNotFound:
		return c.Status(fiber.StatusOK).SendString("nothing created")

	default:
		return c.SendStatus(fiber.StatusInternalServerError)
	}
}
