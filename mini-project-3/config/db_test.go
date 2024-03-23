package config_test

import (
	"fmt"
	"sekolahbeta/mini-project-3/config"
	"testing"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("env not found, using system env")
	}
}

func TestConnection(t *testing.T) {
	Init()
	config.OpenDB("testing")
}
