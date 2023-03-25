package controllers

import (
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
	
	"github.com/gin-gonic/gin"

	"github.com/changdaozheng/headhome-backend/logic"
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

	//Convert io.Reader data type to []byte data type
	bytesData := []byte(reqBod)

	//Create new sos log
	res, err := database.CreateSOSLog(bytesData); 
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{"SOSId": res})
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
func AcceptSOSRequest(c *gin.Context) {
	//Process request body
	type requestBody struct{
		VId string `json:"VId"`
		AuthID string `json:"AuthID"`
		SOSId string `json:"SOSId"`
	}

	var req requestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	//Retrieve sosLog
	sosLog, err := database.FindSOSLog(req.SOSId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "sos record not found"})
		return
	}
	
	//Retrieve care receiver involved
	careReceiver, err := database.ReadCareReceiver(sosLog.CrId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "care receiver not found"})
		return
	}

	//Retrieve requesting volunteer
	volunteer, err := database.ReadVolunteer(req.VId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "volunteer not found"})
		return
	}

	//Authenticate and verify status
	currentTime := time.Now().Unix()
	if volunteer.CertificationStart >= currentTime || volunteer.CertificationEnd <= currentTime {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "volunteer not certified"})
		return
	} else if req.AuthID != careReceiver.AuthID {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "authentication failed"})
		return
	} else if sosLog.Status != "lost" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "care receiver have already received help, thank you!"})
		return
	} else {
		//Declare data to be updated and convert to []byte
		data := map[string]interface{}{
			"Volunteer": req.VId,
			"VolunteerContactNum": volunteer.ContactNum,
			"Status": "guided",
		}
		bytesData, err := json.Marshal(data)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		
		//Update SOS Log 
		err = database.UpdateSOSLog(bytesData, req.SOSId)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		//Retreive care giver information
		careGiver, err := database.ReadCareGiver(careReceiver.CareGiver[0].Id)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "no care giver found"})
		}

		//Retreive route geometry 
		directions, err := logic.RetrieveDirections(fmt.Sprintf("%f,%f", sosLog.StartLocation.Lat, sosLog.StartLocation.Lng), careGiver.Address)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}

		//Send response
		resMsg := map[string]interface{} {
			"CrId": careReceiver.CrId,
			"Name": careReceiver.Name,
			"Address": careReceiver.Address,
			"ContactNum": careReceiver.ContactNum,
			"CgName": careGiver.Name,
			"CgContactNum": careGiver.ContactNum,
			"RouteGeom": directions.OverallPolyline,
		}
		c.IndentedJSON(http.StatusOK, gin.H{"message":resMsg})

	} 
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