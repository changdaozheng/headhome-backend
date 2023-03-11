package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/changdaozheng/headhome-backend/database"
)

//CRUD Functions

//Add new care receiver
func AddCareReceiver(c *gin.Context){
	if err := database.CreateCareReceiver(c); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}

//Get complete list of care receiver data
func GetAllCareReceivers(c *gin.Context){
	result, err := database.ReadAllCareReceivers()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

//Get specific care receiver data
func GetCareReceiver(c *gin.Context){
	id := c.Param("id")
	result, err := database.ReadCareReceiver(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

//Update data for a specific care receiver
func UpdateCareReceiver(c *gin.Context) {
	id := c.Param("id")
	err := database.UpdateCareReceiver(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}

//Delete data of a specific care receiver
func DeleteCareReceiver(c *gin.Context) {
	id := c.Param("id")
	err := database.DeleteCareReceiver(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}


//Business logic 
