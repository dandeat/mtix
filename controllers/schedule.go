package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"mtix.com/new/models"
)

type CreateSchedule struct {
	MovieId   uint            `json:"movie_id" binding:"required"`
	StartDate string          `json:"start_date" binding:"required"`
	EndDate   string          `json:"end_date" binding:"required"`
	Room      int             `json:"room" binding:"required"`
	Price     decimal.Decimal `json:"price" binding:"required"`
}

type UpdateSchedule struct {
	MovieId   uint    `json:"movie_id"`
	StartDate string  `json:"start_date"`
	EndDate   string  `json:"end_date"`
	Room      int     `json:"room"`
	Price     float64 `json:"price"`
}

// GET /Schedule
func GetSchedule(c *gin.Context) {
	var jadwal []models.Schedule
	models.DB.Find(&jadwal)

	c.JSON(http.StatusOK, gin.H{"data": jadwal})
}

// GET /Schedule/:id
func FindSchedule(c *gin.Context) {
	var jadwal models.Schedule
	if err := models.DB.Where("id = ?", c.Param("id")).First(&jadwal).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": jadwal})
}

// POST /Schedule
func CreateScheduleFunc(c *gin.Context) {
	var input CreateSchedule
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	const layout = "0001-01-01T00:00:00Z"
	st, _ := time.Parse(layout, input.StartDate)
	ed, _ := time.Parse(layout, input.EndDate)
	// Create Schedule
	sd := models.Schedule{
		Movie_ID:   input.MovieId,
		Start_Date: st,
		End_Date:   ed,
		Room:       input.Room,
		Price:      input.Price,
	}
	models.DB.Create(&sd)

	c.JSON(http.StatusOK, gin.H{"data": sd})
}

// PATCH /schedule/:id
func UpdateSchedulefunc(c *gin.Context) {
	// Get model if exist
	var jadwal models.Schedule
	if err := models.DB.Where("id = ?", c.Param("id")).First(&jadwal).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateSchedule
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&jadwal).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": jadwal})
}

// DELETE /schedule/:id
func DeleteSchedule(c *gin.Context) {
	// Get model if exist
	var jadwal models.Schedule
	if err := models.DB.Where("id = ?", c.Param("id")).First(&jadwal).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&jadwal)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
