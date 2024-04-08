package controller

import (
	"college/internal/usecase/models"
	"college/internal/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (ctrl *Controller) CreateScreen(c *fiber.Ctx) error {
	var scr models.Screen

	if err := c.BodyParser(&scr); err != nil {
		_ = c.Status(http.StatusBadRequest).SendString(err.Error())
		return nil
	}

	if err := ctrl.uc.CreateScreen(c.Context(), scr); err != nil {
		st, msg := utils.FromError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	_ = c.Send([]byte("successfully created current screen!"))
	c.Status(http.StatusCreated)
	return nil
}

func (ctrl *Controller) UpdateScreen(c *fiber.Ctx) error {
	var scr models.Screen

	scr.ID = c.Query("id")

	if err := c.BodyParser(&scr); err != nil {
		_ = c.Status(http.StatusBadRequest).SendString(err.Error())
		return nil
	}

	if err := ctrl.uc.UpdateScreen(c.Context(), scr); err != nil {
		st, msg := utils.FromError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	_ = c.Send([]byte("successfully updated current screen!"))
	c.Status(http.StatusOK)
	return nil
}

func (ctrl *Controller) DeleteScreen(c *fiber.Ctx) error {
	id := c.Query("id")

	if err := ctrl.uc.DeleteScreen(c.Context(), id); err != nil {
		st, msg := utils.FromError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	_ = c.Send([]byte("successfully deleted current screen!"))
	c.Status(http.StatusOK)
	return nil
}
