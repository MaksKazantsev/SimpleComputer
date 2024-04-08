package controller

import (
	"college/internal/usecase/models"
	"college/internal/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (ctrl *Controller) CreateMouse(c *fiber.Ctx) error {
	var m models.Mouse

	if err := c.BodyParser(&m); err != nil {
		_ = c.Status(http.StatusBadRequest).SendString(err.Error())
		return nil
	}

	if err := ctrl.uc.CreateMouse(c.Context(), m); err != nil {
		st, msg := utils.FromError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	_ = c.Send([]byte("successfully created current mouse!"))
	c.Status(http.StatusCreated)
	return nil
}

func (ctrl *Controller) UpdateMouse(c *fiber.Ctx) error {
	var m models.Mouse

	m.ID = c.Query("id")

	if err := c.BodyParser(&m); err != nil {
		_ = c.Status(http.StatusBadRequest).SendString(err.Error())
		return nil
	}

	if err := ctrl.uc.UpdateMouse(c.Context(), m); err != nil {
		st, msg := utils.FromError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}
	_ = c.Send([]byte("successfully updated current mouse!"))
	_ = c.Status(http.StatusOK)
	return nil
}

func (ctrl *Controller) DeleteMouse(c *fiber.Ctx) error {
	id := c.Query("id")

	if err := ctrl.uc.DeleteMouse(c.Context(), id); err != nil {
		st, msg := utils.FromError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	_ = c.Send([]byte("successfully deleted current mouse!"))
	c.Status(http.StatusOK)
	return nil
}
