package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/changdaozheng/gin-backend/logic/members"
)

func InitRoutes(router *gin.Engine){
	//members
	router.GET("/members", members.GetMembers)
	router.POST("/members", members.AddMember)
}