package utils

import (
	"fmt"
	"sekolahbeta/pertemuan10/config"
	"sekolahbeta/pertemuan10/model"
	"strconv"
	"time"
)

func GetCarsList() ([]model.Car, error) {
	var cars model.Car
	return cars.GetAll(config.Mysql.DB)
}

func GetCarByID(id uint) (model.Car, error) {
	cars := model.Car{
		Model: model.Model{
			ID: id,
		},
	}
	return cars.GetByID(config.Mysql.DB)
}

func InsertCarData(data model.Car) (model.Car, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := data.Create(config.Mysql.DB)

	return data, err
}

func ImportCsvFile(records [][]string) error {
	cars := []model.Car{}

	for i, car := range records {
		if i == 0 {
			continue
		}

		SellingPrice, err := strconv.Atoi(car[14])
		if err != nil {
			SellingPrice = 0
		}
		cars = append(cars, model.Car{
			UUID:         car[0],
			Nama:         car[3],
			Tipe:         car[5],
			Tahun:        car[1],
			Color:        car[10],
			Condition:    car[8],
			SellingPrice: SellingPrice,
		})
	}

	for _, car := range cars {
		car.CreatedAt = time.Now()
		car.UpdatedAt = time.Now()

		res, err := car.GetByUUID(config.Mysql.DB)
		if err != nil {
			if err.Error() != "record not found" {
				return err
			} else {
				err = car.Create(config.Mysql.DB)
				if err != nil {
					return fmt.Errorf("failed to import data, error :%s", err.Error())
				}
			}
		} else {
			car.ID = res.ID
			car.CreatedAt = res.CreatedAt

			err = car.UpdateOneByID(config.Mysql.DB)
			if err != nil {
				return fmt.Errorf("failed to update data, error :%s", err.Error())
			}
		}
	}
	return nil
}
