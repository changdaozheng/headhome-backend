package controllers

import(
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

//WEBSOCKET Request
func Help(c *gin.Context) {
	//Extract request body
	CrId := c.Param("id")
	var req map[string]string
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	//send help message
	if err := fcm.TopicSend(req, CrId); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	//Call handler function
	result, err := logic.RetrieveDirections(req["Start"], req["End"])
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}