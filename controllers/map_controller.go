package controllers

import(
	"fmt"
	"net/http"
	
	"github.com/gin-gonic/gin"

	"github.com/changdaozheng/headhome-backend/logic"
	"github.com/changdaozheng/headhome-backend/fcm"
)
//Maps API Request
func PlanRoute(c *gin.Context){
	//Process request
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	//Call handler function
	result, err := logic.RetrieveDirections(req["Start"].(string), req["End"].(string))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	
	c.IndentedJSON(http.StatusOK, result)
	return
}

//FCM Request
func Help(c *gin.Context) {
	//Extract request body
	CrId := c.Param("id")
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	//send help message
	// convert values to map[string]string
	strMap := make(map[string]string)

	for key, value := range req {
		switch stringValue := value.(type) {
		case string:
			strMap[key] = stringValue
		default:
			strMap[key] = fmt.Sprintf("%v", value)
		}
	}

	if err := fcm.TopicSend(strMap, CrId); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	//Call handler function
	result, err := logic.RetrieveDirections(req["Start"].(string), req["End"].(string))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}