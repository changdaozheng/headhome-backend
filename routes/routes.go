package routes

import (
	"github.com/gin-gonic/gin"
	
	"github.com/changdaozheng/headhome-backend/controllers"
)

func InitRoutes(router *gin.Engine){
	//volunteers
	volunteerR := router.Group("/volunteers")
	volunteerR.GET("", controllers.GetVolunteers)
	volunteerR.GET("/:id", controllers.GetVolunteer)
	volunteerR.POST("", controllers.AddVolunteer)
	volunteerR.PUT("/:id", controllers.UpdateVolunteer)
	volunteerR.DELETE("/:id", controllers.DeleteVolunteer)

}