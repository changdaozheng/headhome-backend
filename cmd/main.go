package main

import (
	"github.com/gin-gonic/gin"

	"github.com/changdaozheng/headhome-backend/routes"
	"github.com/changdaozheng/headhome-backend/database"
)

func main(){
	router := gin.Default()
	database.InitDB()
	routes.InitRoutes(router)
	router.Run("localhost:8080")
	defer database.CloseDB()
}