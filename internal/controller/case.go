package controller

import (
	"college/internal/usecase/models"
	"college/internal/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (ctrl *Controller) CreateCase(c *fiber.Ctx) error {
	var pc models.Case

	if err := c.BodyParser(&pc); err != nil {
		_ = c.Status(http.StatusBadRequest).SendString(err.Error())
		return nil
	}

	if err := ctrl.uc.CreateCase(c.Context(), pc); err != nil {
		st, msg := utils.FromError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	_ = c.Send([]byte("successfully created current pc case!"))
	c.Status(http.StatusCreated)
	return nil
}

func (ctrl *Controller) DeleteCase(c *fiber.Ctx) error {
	id := c.Query("id")

	if err := ctrl.uc.DeleteCase(c.Context(), id); err != nil {
		st, msg := utils.FromError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	_ = c.Send([]byte("successfully deleted current pc case!"))
	c.Status(http.StatusOK)
	return nil
}

func (ctrl *Controller) UpdateCase(c *fiber.Ctx) error {
	var p models.Case

	p.ID = c.Query("id")

	if err := c.BodyParser(&p); err != nil {
		_ = c.Status(http.StatusBadRequest).SendString(err.Error())
		return nil
	}

	if err := ctrl.uc.UpdateCase(c.Context(), p); err != nil {
		st, msg := utils.FromError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	_ = c.Send([]byte("successfully updated current pc case!"))
	c.Status(http.StatusOK)
	return nil
}
