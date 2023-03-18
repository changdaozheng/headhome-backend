package database

import (
	"fmt"
	"errors"
	"reflect"
	"strconv"
	"encoding/json"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	"github.com/changdaozheng/headhome-backend/models"
)

var sosLogRef *firestore.CollectionRef

//Initialise in database.go
func InitSosLog() {
	sosLogRef = Client.Collection("sos_log")
}

//Create new document
func CreateSOSLog(data []byte) (string, error) {
	//Unmarshal data
	var sosLog models.SOSLog
	if err := json.Unmarshal(data, &sosLog); err != nil {
		return "", err
	}

	//Create document with composite id
	sosLogId := sosLog.CrId + strconv.Itoa(int(sosLog.Datetime))
	sosLog.SOSId = sosLogId
	_, err := sosLogRef.Doc(sosLogId).Set(FBCtx, sosLog)
	if err != nil {
		return "", err
	}
	return sosLogId, nil
}

//Read all documents
func ReadAllSOSLogs() ([]models.SOSLog, error) {
	
	var sosLogs []models.SOSLog

	//Read all documents in collection
	iter := sosLogRef.Documents(FBCtx)
	for {
		//Reading individual documents
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var sosLog models.SOSLog
		if err := doc.DataTo(&sosLog); err != nil {
			return nil, err
		}
		
		// Add document into return array
		sosLogs = append(sosLogs, sosLog)
	}
	return sosLogs, nil
}

//Read latest document from specified care receiver with care receiver's id as input
func ReadLatestSOSLog(id string) (models.SOSLog, error) {
	// Firebase query to find latest document
	q := sosLogRef.Where("cr_id", "==", id)
	q = q.OrderBy("datetime", firestore.Desc).Limit(1)

	// Read the only document
	iter := q.Documents(FBCtx)
	doc, err := iter.Next()
	print(doc)
	if err == iterator.Done {
		return models.SOSLog{}, errors.New(fmt.Sprintf("No SOS Logs found for %s", id))
	} 
	if err != nil {
		return models.SOSLog{}, err
	}

	//Return document
	var sosLog models.SOSLog
	if err := doc.DataTo(&sosLog); err != nil {
		return models.SOSLog{}, err
	}
	return sosLog, nil
}

//Read all document from specified care receiver with care receiver's id as input
func FindSOSLog(SOSId string) (models.SOSLog, error) {
	doc, err := sosLogRef.Doc(SOSId).Get(FBCtx)
	if err != nil {
		return models.SOSLog{}, err
	}

	var sosLog models.SOSLog
	if err := doc.DataTo(&sosLog); err != nil {
		return models.SOSLog{}, err
	}
	return sosLog, nil
}

//Update SOS log's information (e.g. volunteers and status)
func UpdateSOSLog(data []byte, id string) (error) {
	
	//Unmarshal data
	var sosLog models.SOSLog
	if err := json.Unmarshal(data, &sosLog); err != nil {
		return err
	}

	//Unpacking all updates value fields
	updates := []firestore.Update{}
    v := reflect.ValueOf(sosLog)
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
	_, err := sosLogRef.Doc(id).Update(FBCtx, updates)
	if err != nil {
		return err
	}
	return nil
}