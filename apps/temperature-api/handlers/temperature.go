package handlers

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"strconv"
	"time"

	"temperature-api/models"

	"github.com/gin-gonic/gin"
)

// TemperatureHandler handles sensor-related requests
type TemperatureHandler struct {
}

// NewTemperatureHandler creates a new TemperatureHandler
func NewTemperatureHandler() *TemperatureHandler {
	return &TemperatureHandler{}
}

// RegisterRoutes registers the sensor routes
func (h *TemperatureHandler) RegisterRoutes(router *gin.RouterGroup) {
	temperature := router.Group("/temperature")
	{
		temperature.GET("/:id", h.GetTemperatureBySensorID)
		temperature.GET("/", h.GetTemperatureByLocation)
	}
}

// GetTemperatureBySensorID handles GET /temperature/:id
func (h *TemperatureHandler) GetTemperatureBySensorID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sensor ID"})
		return
	}
	location := ""
	switch id {
	case 1:
		location = "Living Room"
	case 2:
		location = "Bedroom"
	case 3:
		location = "Kitchen"
	default:
		location = "Unknown"
	}

	c.JSON(http.StatusOK, getTemperatureData(location, fmt.Sprintf("%d", id)))
}

// GetTemperatureByLocation handles GET /temperature
func (h *TemperatureHandler) GetTemperatureByLocation(c *gin.Context) {
	location, _ := c.GetQuery("location")
	if location == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Location is required"})
		return
	}
	sensorID := ""
	if sensorID == "" {
		switch location {
		case "Living Room":
			sensorID = "1"
		case "Bedroom":
			sensorID = "2"
		case "Kitchen":
			sensorID = "3"
		default:
			sensorID = "0"
		}
	}

	c.JSON(http.StatusOK, getTemperatureData(location, sensorID))
}

func getTemperatureData(location, id string) *models.Temperature {
	return &models.Temperature{
		Location:    location,
		SensorID:    id,
		Timestamp:   time.Now(),
		Value:       10 + float64(rand.IntN(100))/10,
		Unit:        "°C",
		Status:      "active",
		Description: "Temperature data for " + location,
	}
}
