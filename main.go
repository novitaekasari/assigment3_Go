package main

import (
	"log"
	"math/rand"
	"time"

	weatherController "github.com/npvitaekasari/assigment3-go/controllers"
	"github.com/npvitaekasari/assigment3-go/models"
) 



func main() {
	models.ConnectDatabase()
	updateData()
}

func updateData() {
	for{
		water := rand.Float64()*100 +11
		wind := rand.Float64()*100 + 1

		var waterStatus string
		if water < 5 {
			waterStatus = "aman"
		} else if water >= 6 && water <= 8 {
			waterStatus = "siaga"
		} else {
			waterStatus = "bahaya"
		}

		var windStatus string
		if wind < 6 {
			windStatus = "aman"
		} else if wind >= 7 && wind <= 15 {
			windStatus = "siaga"
		} else {
			windStatus = "bahaya"
		}

		log.Printf("Water: %.2f m (%s), wind: %.2f m/s (%s)\n", water, waterStatus, wind, windStatus)

		weatherController.Update(water, wind, waterStatus, windStatus)

		time.Sleep(15 * time.Second)

	}
}