package controllers

import (
	"sekolahbeta/restapidb/config"
	"sekolahbeta/restapidb/model"
	"sekolahbeta/restapidb/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func RouteCars(app *fiber.App) {
	carsGroup := app.Group("/cars")
	carsGroup.Get("/", GetCarsList)
	carsGroup.Get("/by-id/:id", GetCarsByID)
	carsGroup.Post("/", PostCar)
}

func GetCarsList(c *fiber.Ctx) error {
	carsData, err := utils.GetCarsList()
	if err != nil {
		logrus.Error("Error on get cars list: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			})
	}

	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    carsData,
			"message": "Success",
		},
	)
}

func GetCarsByID(c *fiber.Ctx) error {
	carId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "ID not valid",
			})
	}

	carData, err := utils.GetCarID(uint(carId))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "ID not found",
				},
			)
		}
		logrus.Error("Error on get car data", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    carData,
			"message": "Success",
		},
	)
}

func PostCar(c *fiber.Ctx) error {
	req := new(model.Car)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": "Body request not valid",
		})
	}

	car := model.Car{
		Nama:  req.Nama,
		Tipe:  req.Tipe,
		Tahun: req.Tahun,
	}

	err := car.Create(config.Mysql.DB)
	if err != nil {
		logrus.Error("Error on get car data", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}

	return c.Status(fiber.StatusCreated).JSON(
		map[string]any{
			"message": "Success",
		},
	)
}
