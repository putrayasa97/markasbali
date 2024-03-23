package config_test

import (
	"fmt"
	"sekolahbeta/demodatabase/config"
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

func TestKoneksi(t *testing.T) {
	Init()

	_, err := config.OpenConn()
	assert.Nil(t, err)
}
