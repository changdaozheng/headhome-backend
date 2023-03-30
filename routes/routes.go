package routes

import (
	"github.com/gin-gonic/gin"
	
	"github.com/changdaozheng/headhome-backend/controllers"
)

func InitRoutes(router *gin.Engine){
	//API health 
	router.HEAD("/", func(c *gin.Context){c.Status(200)})
	router.GET("/", func(c *gin.Context){c.String(200, "API HEALTHY")})

	
	//volunteers
	volunteerR := router.Group("/volunteers")
	volunteerR.GET("", controllers.GetAllVolunteers)
	volunteerR.GET("/:id", controllers.GetVolunteer)
	volunteerR.POST("", controllers.AddVolunteer)
	volunteerR.PUT("/:id", controllers.UpdateVolunteer)
	volunteerR.DELETE("/:id", controllers.DeleteVolunteer)

	//caregiver 
	careGiverR := router.Group("/caregiver")
	careGiverR.GET("", controllers.GetAllCareGivers)
	careGiverR.GET("/:id", controllers.GetCareGiver)
	careGiverR.POST("", controllers.AddCareGiver)
	careGiverR.PUT("/:id", controllers.UpdateCareGiver)
	careGiverR.PUT("/:id/newcr", controllers.NewCareReceiver)
	careGiverR.PUT("/:id/rmcr", controllers.RemoveCareReceiver)
	careGiverR.DELETE("/:id", controllers.DeleteCareGiver)
	

	//carereceiver
	careReceiverR := router.Group("/carereceiver")
	careReceiverR.GET("", controllers.GetAllCareReceivers)
	careReceiverR.GET("/:id", controllers.GetCareReceiver)
	careReceiverR.GET("/contactcg", controllers.ContactCareGiver)
	careReceiverR.POST("/route", controllers.PlanRoute)
	careReceiverR.POST("", controllers.AddCareReceiver)
	careReceiverR.POST("/:id/help", controllers.Help)
	careReceiverR.PUT("/:id", controllers.UpdateCareReceiver)
	careReceiverR.DELETE("/:id", controllers.DeleteCareReceiver)

	//sos
	sosR := router.Group("/sos")
	sosR.GET("", controllers.GetAllSOSLogs)
	sosR.GET("/:id", controllers.GetLatestSOSLog)
	sosR.POST("/", controllers.AddSOSLog)
	sosR.PUT("/accept", controllers.AcceptSOSRequest)
	sosR.PUT("/:id", controllers.UpdateSOSStatus) 

	//travellog
	travelLogR := router.Group("/travellog")
	travelLogR.GET("/:id", controllers.GetLatestTravelLog)
	travelLogR.GET("/:id/all", controllers.GetTravelLog)
	travelLogR.POST("/:id", controllers.AddTravelLog)
}