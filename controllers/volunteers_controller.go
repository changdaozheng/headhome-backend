package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/changdaozheng/headhome-backend/database"
)

func AddVolunteer(c *gin.Context){
	if err := database.CreateVolunteer(c); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}

func GetAllVolunteers(c *gin.Context){
	result, err := database.ReadAllVolunteers()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

func GetVolunteer(c *gin.Context){
	id := c.Param("id")
	result, err := database.ReadVolunteer(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}


func UpdateVolunteer(c *gin.Context) {
	id := c.Param("id")
	err := database.UpdateVolunteer(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}


func DeleteVolunteer(c *gin.Context) {
	id := c.Param("id")
	err := database.DeleteVolunteer(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}