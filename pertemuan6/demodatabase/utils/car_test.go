package utils_test

import (
	"context"
	"fmt"
	"sekolahbeta/demodatabase/config"
	"sekolahbeta/demodatabase/model"
	"sekolahbeta/demodatabase/utils"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("env not found, using system env")
	}
}

func TestCreateDataSuccess(t *testing.T) {
	Init()

	conn, err := config.OpenConn()
	assert.Nil(t, err)

	bdy := model.Car{
		ID:    "1",
		Nama:  "Toyota",
		Tipe:  "Yaris",
		Tahun: "2018",
	}

	err = utils.InsertData(conn, bdy, context.TODO())
	assert.Nil(nil, err)
}

func TestCreateDataFailed(t *testing.T) {
	Init()

	conn, err := config.OpenConn()
	assert.Nil(t, err)

	bdy := model.Car{
		ID:    "1234",
		Nama:  "Toyota",
		Tipe:  "Yaris",
		Tahun: "2018",
	}

	err = utils.InsertData(conn, bdy, context.TODO())
	assert.Nil(nil, err)

	bdy1 := model.Car{
		ID:    "1234",
		Nama:  "Toyota",
		Tipe:  "Yaris",
		Tahun: "2018",
	}

	err1 := utils.InsertData(conn, bdy1, context.TODO())
	assert.NotNil(nil, err1)
}

func TestGetByID(t *testing.T) {
	Init()

	conn, err := config.OpenConn()
	assert.Nil(t, err)

	res, _ := utils.GetByID(conn, "1234", context.TODO())
	assert.Nil(t, err)
	assert.NotEqual(t, model.Car{}, res)
	assert.Equal(t, "1234", res.ID)

	fmt.Println(res)

}
