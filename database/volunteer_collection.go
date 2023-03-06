package database

import (	
	"fmt"
	"google.golang.org/api/iterator"
	"github.com/gin-gonic/gin"
	"cloud.google.com/go/firestore"

	"github.com/changdaozheng/headhome-backend/models"
)

var volunteerRef *firestore.CollectionRef

func InitVolunteers(){
	volunteerRef = Client.Collection("volunteers")
}

func CreateVolunteer(c *gin.Context) (error) {
	var volunteer models.Volunteer
	if err := c.ShouldBindJSON(&volunteer); err != nil {
		return err
	}

	_, err := volunteerRef.Doc(volunteer.VId).Set(FBCtx, volunteer)
	if err != nil {
		return err
	}

	return nil 
}

func ReadVolunteers() ([]models.Volunteer, error) {
	var volunteers []models.Volunteer
	iter := volunteerRef.Documents(FBCtx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var volunteer models.Volunteer
		if err := doc.DataTo(&volunteer); err != nil {
			return nil, err
		}
		volunteers = append(volunteers, volunteer)
	}
	
	return volunteers, nil
}

func ReadVolunteer(id string) (models.Volunteer, error) {
	
	doc, err := volunteerRef.Doc(id).Get(FBCtx)
	if err != nil {
		return models.Volunteer{}, err
	}

	var volunteer models.Volunteer
	if err := doc.DataTo(&volunteer); err != nil {
		return models.Volunteer{}, err
	}
	return volunteer, nil
}

func UpdateVolunteer(c *gin.Context, id string) (error){
	var volunteer models.Volunteer
	if err := c.ShouldBindJSON(&volunteer); err != nil {
		return err
	}
	_, err := volunteerRef.Doc(id).Set(FBCtx, volunteer)
	if err != nil {

		return err
	}
	return nil
}

func DeleteVolunteer(id string) (error) {
	result, err := volunteerRef.Doc(id).Delete(FBCtx)
	fmt.Print(result)
	if err != nil {
		return err
	}
	return nil
}