package controllers

import(
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/changdaozheng/headhome-backend/logic"
)

type routePlanningReqBod struct {
	Start 	string	`json:"Start"`
	End		string  `json:"End"`
}

func PlanRoute(c *gin.Context){
	//Process request
	var req routePlanningReqBod
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	//Call handler function
	result, err := logic.RetrieveDirections(req.Start, req.End)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	
	c.IndentedJSON(http.StatusOK, result)
	return
}