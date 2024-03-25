package main

import (
	"fmt"
	"sekolahbeta/restapidb/config"
	"sekolahbeta/restapidb/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func InitEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("env not found, using system env")
	}
}

func main() {
	InitEnv()
	config.OpenDB()

	app := fiber.New()
	controllers.RouteCars(app)

	err := app.Listen(":3000")
	if err != nil {
		logrus.Fatal("Error on running fiber, ", err.Error())
	}
}
