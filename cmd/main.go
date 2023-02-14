package main

import (
	"github.com/gin-gonic/gin"

	"github.com/changdaozheng/gin-backend/routes"
	"github.com/changdaozheng/gin-backend/database"
)

func main(){
	router := gin.Default()
	database.InitDB()
	routes.InitRoutes(router)
	router.Run("localhost:8080")
}