package controller

import (
	"college/internal/usecase/models"
	"college/internal/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (ctrl *Controller) CreateKeyBoard(c *fiber.Ctx) error {
	var kb models.KeyBoard

	if err := c.BodyParser(&kb); err != nil {
		_ = c.Status(http.StatusBadRequest).SendString(err.Error())
		return nil
	}

	if err := ctrl.uc.CreateKeyboard(c.Context(), kb); err != nil {
		st, msg := utils.FromError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}

	_ = c.Send([]byte("successfully created current keyboard!"))
	c.Status(http.StatusCreated)
	return nil
}

func (ctrl *Controller) UpdateKeyBoard(c *fiber.Ctx) error {
	var kb models.KeyBoard

	kb.ID = c.Query("id")

	if err := c.BodyParser(&kb); err != nil {
		_ = c.Status(http.StatusBadRequest).SendString(err.Error())
		return nil
	}

	if err := ctrl.uc.UpdateKeyBoard(c.Context(), kb); err != nil {
		st, msg := utils.FromError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}
	_ = c.Send([]byte("successfully updated current keyboard!"))
	c.Status(http.StatusOK)
	return nil
}

func (ctrl *Controller) DeleteKeyBoard(c *fiber.Ctx) error {
	id := c.Query("id")

	if err := ctrl.uc.DeleteKeyBoard(c.Context(), id); err != nil {
		st, msg := utils.FromError(err)
		_ = c.Status(st).SendString(msg)
		return nil
	}
	_ = c.Send([]byte("successfully deleted current keyboard!"))
	_ = c.Status(http.StatusOK)
	return nil
}
