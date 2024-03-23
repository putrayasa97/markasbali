package main

import (
	"fmt"
	"sekolahbeta/mini-project-3/config"
	"sekolahbeta/mini-project-3/utils/menu"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("env not found, using system env")
	}
	config.OpenDB("default")
	menu.MainMenu()
}
