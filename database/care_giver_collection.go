package database

import (
	
	"google.golang.org/api/iterator"
	"github.com/gin-gonic/gin"
	"cloud.google.com/go/firestore"

	"github.com/changdaozheng/headhome-backend/models"
)

var careGiverRef *firestore.CollectionRef

func InitCareGiver(){
	careGiverRef = Client.Collection("care_giver")
}

func CreateCareGiver(c *gin.Context) (error){
	var careGiver models.CareGiver
	if err := c.ShouldBindJSON(&careGiver); err != nil {
		return err
	}

	_, err := careGiverRef.Doc(careGiver.CgId).Set(FBCtx, careGiver)
	if err != nil {
		return err
	}
	return nil
}

func ReadCareGivers() ([]models.CareGiver, error) {
	var careGivers []models.CareGiver
	iter := careGiverRef.Documents(FBCtx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var careGiver models.CareGiver
		if err := doc.DataTo(&careGiver); err != nil {
			return nil, err
		}
		careGivers = append(careGivers, careGiver)
	}
	
	return careGivers, nil
}

func ReadCareGiver(id string) (models.CareGiver, error) {
	
	doc, err := careGiverRef.Doc(id).Get(FBCtx)
	if err != nil {
		return models.CareGiver{}, err
	}

	var careGiver models.CareGiver
	if err := doc.DataTo(&careGiver); err != nil {
		return models.CareGiver{}, err
	}
	return careGiver, nil
}

func UpdateCareGiver(c *gin.Context, id string) (error){
	var careGiver models.CareGiver
	if err := c.ShouldBindJSON(&careGiver); err != nil {
		return err
	}
	_, err := careGiverRef.Doc(id).Set(FBCtx, careGiver)
	if err != nil {

		return err
	}
	return nil
}

func DeleteCareGiver(id string) (error) {
	_, err := careGiverRef.Doc(id).Delete(FBCtx)
	if err != nil {
		return err
	}
	return nil
}