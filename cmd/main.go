package main

import (
	"college/internal/app"
	"college/internal/config"
)

func main() {
	app.MustStart(config.MustLoad())
}
