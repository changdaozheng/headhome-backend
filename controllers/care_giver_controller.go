package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/changdaozheng/headhome-backend/database"
)

func AddCareGiver(c *gin.Context){
	if err := database.CreateCareGiver(c); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}

func GetAllCareGivers(c *gin.Context){
	result, err := database.ReadAllCareGivers()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

func GetCareGiver(c *gin.Context){
	id := c.Param("id")
	result, err := database.ReadCareGiver(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}


func UpdateCareGiver(c *gin.Context) {
	id := c.Param("id")
	err := database.UpdateCareGiver(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}


func NewCareReceiver(c *gin.Context) {
	id := c.Param("id")
	err := database.NewCareReceiver(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}

func RemoveCareReceiver(c *gin.Context) {
	id := c.Param("id")
	err := database.RemoveCareReceiver(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}


func DeleteCareGiver(c *gin.Context) {
	id := c.Param("id")
	err := database.DeleteCareGiver(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}