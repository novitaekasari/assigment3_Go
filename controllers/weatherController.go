package weatherController

import (
	"log"

	"github.com/npvitaekasari/assigment3-go/models"
)

func Update(water float64, wind float64, waterStatus string, windStatus string) {
	var weather models.Weather

	err := models.DB.First(&weather).Error;
	
	weather.Water = water
	weather.Wind = wind
	weather.WaterStatus = waterStatus
	weather.WindStatus = windStatus

	if err != nil {
		if models.DB.Create(&weather).RowsAffected == 0 {
			log.Fatal("Tidak dapat mengupdate weather")
		}
	}

	if models.DB.Save(&weather).RowsAffected == 0 {
		log.Fatal("Tidak dapat mengupdate weather")
	}
}