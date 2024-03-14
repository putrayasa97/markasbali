package controllers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sekolahbeta/hacker/model"
	"sync"

	"github.com/gofiber/fiber/v2"
)

func GetCar(c *fiber.Ctx) error {
	fileCsv, err := os.Open("cars_500.csv")
	if err != nil {
		fmt.Println(err)
	}

	defer fileCsv.Close()

	reader := csv.NewReader(fileCsv)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	cars := []model.Car{}

	wg := sync.WaitGroup{}
	channel := make(chan []string)
	channelCar := make(chan model.Car, len(records))
	jml := 5

	for i := 0; i < jml; i++ {
		wg.Add(1)
		go func(channel <-chan []string, channelCar chan model.Car, wg *sync.WaitGroup) {
			for car := range channel {
				channelCar <- model.Car{
					ID:           car[0],
					Year:         car[1],
					Make:         car[2],
					Model:        car[3],
					Trim:         car[4],
					Body:         car[5],
					Transmission: car[6],
					State:        car[7],
					Condition:    car[8],
					Odometer:     car[9],
					Color:        car[10],
					Interior:     car[11],
					Seller:       car[12],
					Mmr:          car[13],
					SellingPrice: car[14],
					SaleDate:     car[15],
				}
			}
			wg.Done()
		}(channel, channelCar, &wg)
	}

	for i, car := range records {
		if i == 0 {
			continue
		}
		channel <- car
	}

	close(channel)
	wg.Wait()
	close(channelCar)

	for car := range channelCar {
		cars = append(cars, car)
	}

	fmt.Println(len(cars))
	encoded, err := json.MarshalIndent(cars, "", "    ")

	if err != nil {
		fmt.Println(err)
	}

	return c.Send(encoded)
}
