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

// import (
// 	"fmt"
// 	"github.com/changdaozheng/headhome-backend/logic"
// )

// func main(){
// 	res, _ := logic.RetrieveDirections("1.344217,103.682791", "1.354070,103.687003")
// 	fmt.Println(res)
// }
