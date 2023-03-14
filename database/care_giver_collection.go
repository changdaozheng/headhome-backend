package database

import (
	"reflect"
	
	"google.golang.org/api/iterator"
	"github.com/gin-gonic/gin"
	"cloud.google.com/go/firestore"

	"github.com/changdaozheng/headhome-backend/models"
)

var careGiverRef *firestore.CollectionRef

//Initialise in database.go
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

func NewCareReceiver(c *gin.Context, id string) (error){
	var newCareReceiver models.Relationship
	if err := c.ShouldBindJSON(&newCareReceiver); err != nil {
		return err
	}

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

func RemoveCareReceiver(c *gin.Context, id string) (error){
	var careReceiver models.Relationship
	if err := c.ShouldBindJSON(&careReceiver); err != nil {
		return err
	}

	update := []firestore.Update {
		{
			Path:  "care_receiver",
			Value: firestore.ArrayRemove(careReceiver.Id),
		},
	}

	_, err := careGiverRef.Doc(id).Update(FBCtx, update)
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