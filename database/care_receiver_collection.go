package database

import (
	"reflect"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	"github.com/changdaozheng/headhome-backend/logic"
	"github.com/changdaozheng/headhome-backend/models"
)

var careReceiverRef *firestore.CollectionRef

func InitCareReceiver(){
	careReceiverRef = Client.Collection("care_receiver")
}

func CreateCareReceiver(data []byte) (error){
	//Unmarshal data
	var careReceiver models.CareReceiver
	if err := json.Unmarshal(data, &careReceiver); err != nil {
		return err
	}
	
	//Generating Auth ID
	careReceiver.AuthID = logic.RandStr(10)

	//Create care receiver
	_, err := careReceiverRef.Doc(careReceiver.CrId).Set(FBCtx, careReceiver)
	if err != nil {
		return err
	}
	return nil
}

func ReadAllCareReceivers() ([]models.CareReceiver, error) {
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

//Update care receiver details (except care giver modifications e.g. add and remove)
func UpdateCareReceiver(c *gin.Context, id string) (error){
	var careReceiver models.CareReceiver
	if err := c.ShouldBindJSON(&careReceiver); err != nil {
		return err
	}

	//remove care giver modification during normal update
	careReceiver.CareGiver = []models.Relationship{}


	updates := []firestore.Update{}
    v := reflect.ValueOf(careReceiver)
    for i := 0; i < v.NumField(); i++ {
        field := v.Type().Field(i)
        value := v.Field(i)
        if value.IsZero() {
            continue
        }
        updates = append(updates, firestore.Update{
            Path:  field.Tag.Get("firestore"),
            Value: value.Interface(),
        })
    }
	
	_, err := careReceiverRef.Doc(id).Update(FBCtx, updates)
	if err != nil {
		return err
	}
	return nil
}

//Alter care giver (only allow 1 care giver per care receier)
func ChangeCareGiver(newCg []models.Relationship, id string) (error) {
	update := []firestore.Update{
		{
			Path: "care_giver",
			Value: newCg,
		},
	}

	if _, err := careReceiverRef.Doc(id).Update(FBCtx, update); err != nil {
		return err
	}
	return nil

}

//Delete care receiver
func DeleteCareReceiver(id string) (error) {
	_, err := careReceiverRef.Doc(id).Delete(FBCtx)
	if err != nil {
		return err
	}
	return nil
}