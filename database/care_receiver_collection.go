package database

import (
	
	"google.golang.org/api/iterator"
	"github.com/gin-gonic/gin"
	"cloud.google.com/go/firestore"

	"github.com/changdaozheng/headhome-backend/models"
)

var careReceiverRef *firestore.CollectionRef

func InitCareReceiver(){
	careReceiverRef = Client.Collection("care_receiver")
}

func CreateCareReceiver(c *gin.Context) (error){
	var careReceiver models.CareReceiver
	if err := c.ShouldBindJSON(&careReceiver); err != nil {
		return err
	}

	_, err := careReceiverRef.Doc(careReceiver.CrId).Set(FBCtx, careReceiver)
	if err != nil {
		return err
	}
	return nil
}

func ReadCareReceivers() ([]models.CareReceiver, error) {
	var careReceivers []models.CareReceiver
	iter := careReceiverRef.Documents(FBCtx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var careReceiver models.CareReceiver
		if err := doc.DataTo(&careReceiver); err != nil {
			return nil, err
		}
		careReceivers = append(careReceivers, careReceiver)
	}
	
	return careReceivers, nil
}

func ReadCareReceiver(id string) (models.CareReceiver, error) {
	
	doc, err := careReceiverRef.Doc(id).Get(FBCtx)
	if err != nil {
		return models.CareReceiver{}, err
	}

	var careReceiver models.CareReceiver
	if err := doc.DataTo(&careReceiver); err != nil {
		return models.CareReceiver{}, err
	}
	return careReceiver, nil
}

func UpdateCareReceiver(c *gin.Context, id string) (error){
	var careReceiver models.CareReceiver
	if err := c.ShouldBindJSON(&careReceiver); err != nil {
		return err
	}
	_, err := careReceiverRef.Doc(id).Set(FBCtx, careReceiver)
	if err != nil {

		return err
	}
	return nil
}

func DeleteCareReceiver(id string) (error) {
	_, err := careReceiverRef.Doc(id).Delete(FBCtx)
	if err != nil {
		return err
	}
	return nil
}