package controllers

import(
	"net/http"
	
	"github.com/gin-gonic/gin"

	"github.com/changdaozheng/headhome-backend/logic"
	"github.com/changdaozheng/headhome-backend/websocket"
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
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	//send help message
	if err := websocket.Send(req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()} )
	}

	//Call handler function
	result, err := logic.RetrieveDirections(req["Start"].(string), req["End"].(string))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}