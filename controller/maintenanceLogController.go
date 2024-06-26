package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sokungz01/cpe241-project-backend/domain"
)

type MaintenanceLogController interface {
	GetAllmaintenanceLog(c *fiber.Ctx) error
	GetMaintenanceLogByID(c *fiber.Ctx) error
	GetMaintenanceLogByMachineID(c *fiber.Ctx) error
	GetMaintenanceLogByStaffID(c *fiber.Ctx) error
	CreatemaintenanceLog(c *fiber.Ctx) error
	UpdateMaintenanceLogStatus(c *fiber.Ctx) error
}

type maintenanceLogcontroller struct {
	usecase domain.MaintenanceLogUsecase
}

func NewMaintenanceLogController(usecase domain.MaintenanceLogUsecase) MaintenanceLogController {
	return &maintenanceLogcontroller{usecase: usecase}
}

func (u *maintenanceLogcontroller) GetAllmaintenanceLog(c *fiber.Ctx) error {
	response, err := u.usecase.GetAllmaintenanceLog()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (u *maintenanceLogcontroller) GetMaintenanceLogByMachineID(c *fiber.Ctx) error {
	id, parserr := strconv.Atoi(c.Params("id"))
	if parserr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("maintenanceLog: parse err")
	}
	response, err := u.usecase.GetMaintenanceLogByMachineID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (u *maintenanceLogcontroller) GetMaintenanceLogByStaffID(c *fiber.Ctx) error {
	id, parserr := strconv.Atoi(c.Params("id"))
	if parserr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("maintenanceLog: parse err")
	}
	response, err := u.usecase.GetMaintenanceLogByStaffID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (u *maintenanceLogcontroller) GetMaintenanceLogByID(c *fiber.Ctx) error {
	id, parserr := strconv.Atoi(c.Params("id"))
	if parserr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("maintenanceLog: parse err")
	}
	response, err := u.usecase.GetMaintenanceLogByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (u *maintenanceLogcontroller) CreatemaintenanceLog(c *fiber.Ctx) error {
	newLog := new(domain.MaintenanceLog)
	err := c.BodyParser(newLog)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	response, createErr := u.usecase.CreatemaintenanceLog(newLog)
	if createErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(createErr.Error())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (u *maintenanceLogcontroller) UpdateMaintenanceLogStatus(c *fiber.Ctx) error {
	input := new(domain.MaintenanceLog)
	id, convserr := strconv.Atoi(c.Params("id"))
	if convserr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(convserr.Error())
	}
	parseerr := c.BodyParser(input)
	if parseerr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(parseerr.Error())
	}
	updateErr := u.usecase.UpdateMaintenanceLogStatus(id, input.StatusID)
	if updateErr != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(updateErr.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}
