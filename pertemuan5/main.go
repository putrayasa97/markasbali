package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	// alias
	fiberConfig "sekolahbeta/hacker/config/fiber"
	"sekolahbeta/hacker/controllers"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Get("/pesanan", controllers.GetPesanan)
	app.Get("/car", controllers.GetCar)

	listenAddress := fmt.Sprintf("%s:%s", fiberConfig.GetFiberHttpHost(), fiberConfig.GetFiberHttpPort())
	err := app.Listen(listenAddress)
	if err != nil {
		panic(err)
	}
}
