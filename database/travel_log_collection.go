package database

import (
	"fmt"
	"errors"
	"reflect"
	"strconv"
	"encoding/json"

	"google.golang.org/api/iterator"
	"cloud.google.com/go/firestore"

	"github.com/changdaozheng/headhome-backend/models"
)

var travelLogRef *firestore.CollectionRef

//Initialise in database.go
func InitTravelLog() {
	travelLogRef = Client.Collection("travel_log")
}

//Create new document
func CreateTravelLog(data []byte) (string, error) {
	//Unmarshal data
	var travelLog models.TravelLog
	if err := json.Unmarshal(data, &travelLog); err != nil {
		return "", err
	}
	
	//Create document with composite id
	travelLogId := travelLog.CrId + strconv.Itoa(int(travelLog.Datetime))
	travelLog.TravelLogId = travelLogId
	_, err := travelLogRef.Doc(travelLogId).Set(FBCtx, travelLog)
	if err != nil {
		return "", err
	}
	
	//Check last at home 
	q := travelLogRef.Where("cr_id", "==", travelLog.CrId).Where("status", "==", "home").OrderBy("datetime", firestore.Desc).Limit(1)

	
	iter := q.Documents(FBCtx)
	doc, err := iter.Next()

	if err == iterator.Done {
		return "", errors.New(fmt.Sprintf("%s has no was not at home", travelLog.CrId))
	} 
	if err != nil {
		return "" , err
	}

	//Return document
	if err := doc.DataTo(&travelLog); err != nil {
		return "", err
	}
	return strconv.Itoa(int(travelLog.Datetime)), nil

}

//Read all documents
func ReadAllTravelLogs() ([]models.TravelLog, error) {
	
	var travelLogs []models.TravelLog

	//Read all documents in collection
	iter := travelLogRef.Documents(FBCtx)
	for {
		//Reading individual documents
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var travelLog models.TravelLog
		if err := doc.DataTo(&travelLog); err != nil {
			return nil, err
		}
		
		// Add document into return array
		travelLogs = append(travelLogs, travelLog)
	}
	return travelLogs, nil
}

//Read all documents from specified care receiver
func ReadTravelLog(id string) ([]models.TravelLog, error) {
	//Firebase query to find all documents that belongs to a care receiver
	q := travelLogRef.Where("cr_id", "==", id)
	q = q.OrderBy("datetime", firestore.Desc)

	//Iterate through all documents and return as slice
	var travelLogs []models.TravelLog
	iter := q.Documents(FBCtx)
	for {
		//Reading individual documents
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var travelLog models.TravelLog
		if err := doc.DataTo(&travelLog); err != nil {
			return nil, err
		}
		
		// Add document into return array
		travelLogs = append(travelLogs, travelLog)
	}
	return travelLogs, nil
}

//Read latest document from specified care receiver
func ReadLatestTravelLog(id string) (models.TravelLog, error) {
	// Firebase query to find latest document
	q := travelLogRef.Where("cr_id", "==", id)
	q = q.OrderBy("datetime", firestore.Desc).Limit(1)

	// Read the only document
	iter := q.Documents(FBCtx)
	doc, err := iter.Next()
	if err == iterator.Done {
		return models.TravelLog{}, errors.New(fmt.Sprintf("No SOS Log found for %s", id))
	} 
	if err != nil {
		return models.TravelLog{}, err
	}

	//Return document
	var travelLog models.TravelLog
	if err := doc.DataTo(&travelLog); err != nil {
		return models.TravelLog{}, err
	}
	return travelLog, nil
}

//Update SOS log's information (e.g. volunteers and status)
func UpdateTravelLog(data []byte, id string) (error) {
	
	//Unmarshal data
	var travelLog models.TravelLog
	err := json.Unmarshal(data, &travelLog)
	if err != nil {
		return err
	}
	
	//Unpacking all updates value fields
	updates := []firestore.Update{}
    v := reflect.ValueOf(travelLog)
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

	//Update
	_, err = travelLogRef.Doc(id).Update(FBCtx, updates)
	if err != nil {
		return err
	}
	return nil
}