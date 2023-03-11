package controllers

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/gin-gonic/gin"

	"github.com/changdaozheng/headhome-backend/database"
)

//Add new sos log
func AddSOSLog(c *gin.Context) {
	//Extract request body 
	reqBod, err := ioutil.ReadAll(c.Request.Body)
    if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      	return 
    }

	//convert io.Reader data type to []byte data type
	bytesData := []byte(reqBod)

	//Create new sos log
	if err := database.CreateSOSLog(bytesData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{"message":"successful"})
}

//Read all sos logs 
func GetAllSOSLogs(c *gin.Context) {
	result, err := database.ReadAllSOSLogs()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

//Read latest sos log from specified care receiver 
func GetLatestSOSLog(c *gin.Context) {
	id := c.Param("id")
	result, err := database.ReadLatestSOSLog(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

//Update sos log
type AcceptRequestBody struct{
	VId string `json:"VId"`
	SOSId string `json:"SOSId"`
}

func AcceptSOSRequest(c *gin.Context) {
	//Process request body
	var req *AcceptRequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	//Declare data to be updated and convert to []byte
	data := map[string]interface{}{
		"volunteer": req.VId,
		"status": "guided",
	}
	bytesData, err := json.Marshal(data)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	
	//Update
	err = database.UpdateSOSLog(bytesData, req.SOSId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}

//Update status 
func UpdateSOSStatus(c *gin.Context){
	//Extract information for request body
	id := c.Param("id")
	reqBod, err := ioutil.ReadAll(c.Request.Body)
    if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      	return 
    }

	//convert io.Reader data type to []byte data type
	bytesData := []byte(reqBod)
	err = database.UpdateSOSLog(bytesData, id) 
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"successful"})
}