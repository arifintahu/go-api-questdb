package controllers

import (
	"go-api-questdb/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateTrackerInput struct {
	VehicleId int    `json:"vehicleId"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func CreateTracker(c *gin.Context) {
	var input CreateTrackerInput
	if err:= c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	tracker := models.Tracker{
		Timestamp: time.Now().UTC(),
		VehicleId: input.VehicleId,
		Latitude: input.Latitude,
		Longitude: input.Longitude,
	}

	models.DB.Create(&tracker)

	c.JSON(http.StatusOK, gin.H{"data": tracker})
}

func GetTrackers(c *gin.Context) {
	var trackers []models.Tracker
	models.DB.Find(&trackers)

	c.JSON(http.StatusOK, gin.H{"data": trackers})
}
