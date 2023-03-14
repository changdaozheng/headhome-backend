package controllers

import (
	"net/http"
	"io/ioutil"
	"github.com/gin-gonic/gin"

	"github.com/changdaozheng/headhome-backend/database"
)

//CRUD Functions

//Add new care receiver
func AddCareReceiver(c *gin.Context){
	//Extract request body 
	reqBod, err := ioutil.ReadAll(c.Request.Body)
    if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      	return 
    }

	//Convert io.Reader data type to []byte data type
	bytesData := []byte(reqBod)

	//Create
	if err := database.CreateCareReceiver(bytesData); err != nil {
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

//Return contact number of care giver
func ContactCareGiver(c *gin.Context){
	//Process request body
	type requestBody struct {
		CrId string `json:"CrId"`
		CgId string `json:"CgId"`
	}

	var req requestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest,  gin.H{"error": err.Error()})
		return
	}

	//Retrieve care receiver
	careReceiver, err := database.ReadCareReceiver(req.CrId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "care receiver not found"})
		return
	}

	//Check access permission
	for _, cg := range careReceiver.CareGiver {
		if (cg.Id == req.CgId){
			//Retrieve care giver infromation
			careGiver, err := database.ReadCareGiver(req.CgId)
			if err != nil {
				c.IndentedJSON(http.StatusNotFound, gin.H{"error": "care giver not found"})
				return
			}

			//Send response message
			resMsg := map[string]interface{} {
				"CgContactNum": careGiver.ContactNum,
			}
			c.IndentedJSON(http.StatusOK, resMsg)
			return
		} 
	}
	//None of the linked care givers match requested care giver
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "care giver does not match"})
	return
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