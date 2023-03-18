package controllers

import (
	"errors"
	"net/http"
	
	"github.com/gin-gonic/gin"

	"github.com/changdaozheng/headhome-backend/models"
	"github.com/changdaozheng/headhome-backend/database"
)

//Add new care giver
func AddCareGiver(c *gin.Context){
	if err := database.CreateCareGiver(c); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}

//Read all care giver documents in the system
func GetAllCareGivers(c *gin.Context){
	result, err := database.ReadAllCareGivers()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

//Read specified caregiver
func GetCareGiver(c *gin.Context){
	id := c.Param("id")
	result, err := database.ReadCareGiver(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

//Update specified caregiver details (except care receiver list)
func UpdateCareGiver(c *gin.Context) {
	id := c.Param("id")
	err := database.UpdateCareGiver(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}

//Add new care receiver to care giver's array
func NewCareReceiver(c *gin.Context) {
	cgId := c.Param("id")

	type reqBod struct {
		CrId			string	`json:"CrId"`
		AuthId			string	`json:"AuthId"`
		Relationship	string	`json:"Relationship"`
	}
	
	var req reqBod
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Authenticate 
	careReceiver, err := database.ReadCareReceiver(req.CrId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if careReceiver.AuthID != req.AuthId {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": errors.New("authentication failed")})
		return
	}	

	//Add care receiver to care giver docs	
	newCareReceiver := models.Relationship{
		Id: req.CrId,
		Relationship: req.Relationship,
	}
	
	if err := database.NewCareReceiver(newCareReceiver, cgId); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	//Change care giver in care receiver docs
	newCareGiver := []models.Relationship {
		{
			Id: cgId,
			Relationship: req.Relationship,
		},
	}

	if err := database.ChangeCareGiver(newCareGiver, req.CrId); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}

//Remove care receiver from care giver's array
func RemoveCareReceiver(c *gin.Context) {
	cgId := c.Param("id")

	//Extract request information
	type reqBod struct {
		CrId	string `json:"CrId"`
	}
	var req reqBod
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Remove care receiver from care giver docs
	if err := database.RemoveCareReceiver(cgId, req.CrId); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	//Remove care giver from care receiver docs
	if err := database.ChangeCareGiver([]models.Relationship{}, req.CrId); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}

//Remove care giver
func DeleteCareGiver(c *gin.Context) {
	id := c.Param("id")
	err := database.DeleteCareGiver(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}