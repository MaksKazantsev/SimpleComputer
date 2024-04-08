package app

import (
	"college/internal/config"
	"college/internal/controller"
	"college/internal/db/postgres"
	"college/internal/routes"
	"college/internal/usecase"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func MustStart(cfg *config.Config) {

	repo := postgres.NewRepository(postgres.MustConnect())

	uc := usecase.NewUseCase(repo)

	ctrl := controller.NewController(uc)

	app := fiber.New()

	routes.SetupRoutes(app, ctrl)

	if err := app.Listen(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		panic("Failed to run app" + err.Error())
	}

}
