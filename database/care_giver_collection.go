package database

import (
	"fmt"
	"reflect"
	
	"github.com/gin-gonic/gin"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	"github.com/changdaozheng/headhome-backend/models"
)

var careGiverRef *firestore.CollectionRef

//Initialise in database.go
func InitCareGiver(){
	careGiverRef = Client.Collection("care_giver")
}

//Create new document
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

//Read all documents
func ReadAllCareGivers() ([]models.CareGiver, error) {
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

//Read specific document
func ReadCareGiver(id string) (models.CareGiver, error) {
	
	fmt.Println(FBCtx)
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

//Update care giver details (except care giver modifications e.g. add and remove)
func UpdateCareGiver(c *gin.Context, id string) (error){
	var careGiver models.CareGiver
	if err := c.ShouldBindJSON(&careGiver); err != nil {
		return err
	}

	//remove care receiver modification during normal update
	careGiver.CareReceiver = []models.Relationship{}

	updates := []firestore.Update{}
    v := reflect.ValueOf(careGiver)
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

	_, err := careGiverRef.Doc(id).Update(FBCtx, updates)
	if err != nil {
		return err
	}
	return nil
}

//Add new care receiver to specified care giver

func NewCareReceiver(newCareReceiver models.Relationship, id string) (error){
	update := []firestore.Update {
		{
			Path:  "care_receiver",
			Value: firestore.ArrayUnion(newCareReceiver),
		},
	}

	_, err := careGiverRef.Doc(id).Update(FBCtx, update)
	if err != nil {
		return err
	}
	return nil
}

//Remove care receiver from specified care giver
func RemoveCareReceiver(CgId string, CrId string) (error){

	//ArrayRemove not available in go; Using manual update

	careGiver, err := ReadCareGiver(CgId)
	if err != nil {
		return err
	}

	careReceivers := careGiver.CareReceiver
	for i, cr := range careReceivers {
		if cr.Id == CrId {
			careReceivers = append(careReceivers[:i], careReceivers[i+1:]...)
		}
	}

	update := []firestore.Update {
		firestore.Update{
			Path:  "care_receiver",
			Value: careReceivers,
		},
	}

	
	if _, err := careGiverRef.Doc(CgId).Update(FBCtx, update); err != nil {
		return err
	}
	return nil
}

//Delete care receiver
func DeleteCareGiver(id string) (error) {
	_, err := careGiverRef.Doc(id).Delete(FBCtx)
	if err != nil {
		return err
	}
	return nil
}