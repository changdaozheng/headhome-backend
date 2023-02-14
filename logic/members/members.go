package members

import (
	"net/http"
	
	"google.golang.org/api/iterator"
	"github.com/gin-gonic/gin"
	
	"github.com/changdaozheng/gin-backend/database"
	"github.com/changdaozheng/gin-backend/models"
)


func GetMembers(c *gin.Context){
	
	var members []models.MemberStruct
	iter := database.Members.Documents(database.Ctx)
	for {
			doc, err := iter.Next()
			if err == iterator.Done {
					break
			}
			if err != nil {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			}
			
			var member models.MemberStruct
			if err := doc.DataTo(&member); err != nil {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
			}

			members = append(members, member)
	}
	c.IndentedJSON(http.StatusOK, members)
}

func AddMember(c *gin.Context){
	//Parse POST request
	var member models.MemberStruct
	if err := c.ShouldBindJSON(&member); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
	}
	
	_, _, err := database.Members.Add(database.Ctx, member)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message":"success"})
}