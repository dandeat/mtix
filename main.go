package main

import (
	"mtix.com/new/controllers"
	"mtix.com/new/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	/// Routes
	r.GET("/schedule", controllers.GetSchedule)
	r.GET("/schedule/:id", controllers.FindSchedule)
	r.POST("/schedule", controllers.CreateScheduleFunc)
	r.PATCH("/schedule/:id", controllers.UpdateSchedulefunc)
	r.DELETE("/schedule/:id", controllers.DeleteSchedule)

	// Run the server
	r.Run()
}
