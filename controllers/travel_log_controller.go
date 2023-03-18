package controllers

import (
	"net/http"
	"io/ioutil"
	
	"github.com/gin-gonic/gin"

	"github.com/changdaozheng/headhome-backend/database"
)

//Create new travel log
func AddTravelLog(c *gin.Context) {
	//Extract request body 
	reqBod, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	//convert to bytes
	data := []byte(reqBod)

	//Create 
	
	lastHome, err := database.CreateTravelLog(data)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{"LastHome": lastHome})

}

//Read all travel logs of specified care receiver
func GetTravelLog(c *gin.Context) {
	id := c.Param("id")
	result, err := database.ReadTravelLog(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

//Read latest travel log of specified care receiver
func GetLatestTravelLog(c *gin.Context) {
	id := c.Param("id")
	result, err := database.ReadLatestTravelLog(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

