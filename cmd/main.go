package main

import (
	"github.com/gin-gonic/gin"

	"github.com/changdaozheng/headhome-backend/database"
	"github.com/changdaozheng/headhome-backend/routes"
)


func main(){
	router := gin.Default()
	
	database.InitDB()
	routes.InitRoutes(router)	
	
	router.Run("0.0.0.0:8080")
	defer database.CloseDB()
}