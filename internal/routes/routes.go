package routes

import (
	"college/internal/controller"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func SetupRoutes(a *fiber.App, ctrl *controller.Controller) {

	a.Get("/health", func(c *fiber.Ctx) error {
		c.Status(http.StatusOK)
		return nil
	})

	get := a.Group("/get")
	get.Get("/device", ctrl.GetDevice)
	get.Get("/devices", ctrl.GetAllDevices)

	mouse := a.Group("/mouse")
	mouse.Post("/create", ctrl.CreateMouse)
	mouse.Delete("/delete", ctrl.DeleteMouse)
	mouse.Put("/update", ctrl.UpdateMouse)

	keyboard := a.Group("/keyboard")
	keyboard.Post("/create", ctrl.CreateKeyBoard)
	keyboard.Delete("/delete", ctrl.DeleteKeyBoard)
	keyboard.Put("/update", ctrl.UpdateKeyBoard)

	screen := a.Group("/screen")
	screen.Post("/create", ctrl.CreateScreen)
	screen.Delete("/delete", ctrl.DeleteScreen)
	screen.Put("/update", ctrl.UpdateScreen)

	pc := a.Group("/pc")
	pc.Post("/create", ctrl.CreateCase)
	pc.Delete("/delete", ctrl.DeleteCase)
	pc.Put("/update", ctrl.UpdateCase)
}
