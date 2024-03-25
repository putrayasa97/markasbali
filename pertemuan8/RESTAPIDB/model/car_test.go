package model_test

import (
	"fmt"
	"sekolahbeta/restapidb/config"
	"sekolahbeta/restapidb/model"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("env not found, using system env")
	}
	config.OpenDB()
}

func TestCreateCar(t *testing.T) {
	Init()

	carData := model.Car{
		Nama:  "toyota",
		Tipe:  "crown",
		Tahun: "1998",
	}

	err := carData.Create(config.Mysql.DB)
	assert.Nil(t, err)

	fmt.Println(carData.ID)
}

func TestGetByID(t *testing.T) {
	Init()

	carData := model.Car{
		Model: model.Model{
			ID: 1,
		},
	}

	car, err := carData.GetByID(config.Mysql.DB)
	assert.Nil(t, err)
	fmt.Println(car)
}

func TestGetAll(t *testing.T) {
	Init()

	carData := model.Car{
		Nama:  "Avanza",
		Tipe:  "G",
		Tahun: "2010",
	}

	err := carData.Create(config.Mysql.DB)
	assert.Nil(t, err)

	cars, err := carData.GetAll(config.Mysql.DB)

	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(cars), 1)
	fmt.Println(cars)
}

func TestUpdateOneByID(t *testing.T) {
	Init()

	carData := model.Car{
		Nama:  "toyota",
		Tipe:  "crown",
		Tahun: "1998",
	}

	err := carData.Create(config.Mysql.DB)
	assert.Nil(t, err)

	carData = model.Car{
		Model: model.Model{
			ID: carData.ID,
		},
		Nama:  "Avanzas2",
		Tipe:  "Gs2",
		Tahun: "2015",
	}

	err = carData.UpdateOneByID(config.Mysql.DB)
	assert.Nil(t, err)

	car, err := carData.GetByID(config.Mysql.DB)

	assert.Nil(t, err)
	assert.Equal(t, carData.Nama, car.Nama)
}

func TestDeleteByID(t *testing.T) {
	Init()

	carData := model.Car{
		Nama:  "toyota",
		Tipe:  "crown",
		Tahun: "1998",
	}

	err := carData.Create(config.Mysql.DB)
	assert.Nil(t, err)

	car, err := carData.GetByID(config.Mysql.DB)
	assert.Nil(t, err)
	assert.Equal(t, carData.ID, car.ID)

	err = carData.DeleteByID(config.Mysql.DB)
	assert.Nil(t, err)

	car, err = carData.GetByID(config.Mysql.DB)
	assert.NotNil(t, err)
	assert.Equal(t, model.Car{}, car)

}
