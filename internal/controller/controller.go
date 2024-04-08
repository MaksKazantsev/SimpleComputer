package controller

import (
	"college/internal/usecase"
	"college/internal/usecase/models"
	"college/internal/utils"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Controller struct {
	uc *usecase.UseCase
}

func NewController(uc *usecase.UseCase) *Controller {
	return &Controller{
		uc: uc,
	}
}

const (
	screen   string = "screen"
	mouse    string = "mouse"
	keyboard string = "keyboard"
	pc       string = "pc"
)

func (ctrl *Controller) GetDevice(c *fiber.Ctx) error {

	switch ch := c.Query("type"); ch {
	case screen:
		filter := models.GetScreenReq{
			ID:         c.Query("id"),
			RoomID:     c.Query("room"),
			BuildingID: c.Query("building"),
			Size:       c.Query("size"),
			Herz:       c.Query("herz"),
			Matrix:     c.Query("matrix"),
		}

		scr, err := ctrl.uc.GetScreen(c.Context(), filter)
		if err != nil {
			st, msg := utils.FromError(err)
			_ = c.Status(st).SendString(msg)
			return nil
		}

		b, err := json.Marshal(scr)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return nil
		}

		_ = c.Send(b)
		c.Status(http.StatusOK)
	case mouse:
		filter := models.GetMouseReq{
			ID:         c.Query("id"),
			RoomID:     c.Query("room"),
			BuildingID: c.Query("building"),
			DPI:        c.Query("dpi"),
			Material:   c.Query("material"),
			Size:       c.Query("size"),
			Sensor:     c.Query("sensor"),
		}

		m, err := ctrl.uc.GetMouse(c.Context(), filter)
		if err != nil {
			st, msg := utils.FromError(err)
			_ = c.Status(st).SendString(msg)
			return nil
		}

		b, err := json.Marshal(m)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return nil
		}

		_ = c.Send(b)
		c.Status(http.StatusOK)

	case keyboard:
		filter := models.GetKeyBoardReq{
			ID:         c.Query("id"),
			RoomID:     c.Query("room"),
			BuildingID: c.Query("building"),
			Material:   c.Query("material"),
			Size:       c.Query("size"),
			Switches:   c.Query("switches"),
			KeyType:    c.Query("key_type"),
		}

		kb, err := ctrl.uc.GetKeyboard(c.Context(), filter)
		if err != nil {
			st, msg := utils.FromError(err)
			_ = c.Status(st).SendString(msg)
			return nil
		}
		b, err := json.Marshal(kb)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return nil
		}

		_ = c.Send(b)
		c.Status(http.StatusOK)

	case pc:
		filter := models.GetCaseReq{
			ID:            c.Query("id"),
			RoomID:        c.Query("room"),
			BuildingID:    c.Query("building"),
			Color:         c.Query("color"),
			Size:          c.Query("size"),
			Compatibility: c.Query("comp"),
		}

		pcCase, err := ctrl.uc.GetCase(c.Context(), filter)
		if err != nil {
			st, msg := utils.FromError(err)
			_ = c.Status(st).SendString(msg)
			return nil
		}

		b, err := json.Marshal(pcCase)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return nil
		}

		_ = c.Send(b)
		c.Status(http.StatusOK)
	}
	return nil
}

func (ctrl *Controller) GetAllDevices(c *fiber.Ctx) error {
	roomID := c.Query("room")
	buildingID := c.Query("building")

	switch ch := c.Query("type"); ch {
	case screen:
		scs, err := ctrl.uc.GetAllScreen(c.Context(), roomID, buildingID)
		if err != nil {
			st, msg := utils.FromError(err)
			_ = c.Status(st).SendString(msg)
			return nil
		}

		b, err := json.Marshal(scs)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return nil
		}
		_ = c.Send(b)
	case mouse:
		mice, err := ctrl.uc.GetAllMouse(c.Context(), roomID, buildingID)
		if err != nil {
			st, msg := utils.FromError(err)
			_ = c.Status(st).SendString(msg)
			return nil
		}

		b, err := json.Marshal(mice)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return nil
		}
		_ = c.Send(b)
	case keyboard:
		keyboards, err := ctrl.uc.GetAllKeyBoard(c.Context(), roomID, buildingID)
		if err != nil {
			st, msg := utils.FromError(err)
			_ = c.Status(st).SendString(msg)
			return nil
		}

		b, err := json.Marshal(keyboards)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return nil
		}
		_ = c.Send(b)
	case pc:
		pcCases, err := ctrl.uc.GetAllCase(c.Context(), roomID, buildingID)
		if err != nil {
			st, msg := utils.FromError(err)
			_ = c.Status(st).SendString(msg)
			return nil
		}

		b, err := json.Marshal(pcCases)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return nil
		}

		_ = c.Send(b)
		c.Status(http.StatusOK)

	}
	return nil
}
