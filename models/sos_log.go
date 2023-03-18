package models

type SOSLog struct {
	CrId 		string 	`json:"CrId" firestore:"cr_id"`
	Datetime 	int64 	`json:"Datetime" firestore:"datetime"`
	StartLocation struct{
		Lat float64 `json:"Lat" firestore:"lat"`
		Lng float64 `json:"Lng" firestore:"lng"`
	} `json:"StartLocation" firestore:"start_location"`
	Status 		string 	`json:"Status" firestore:"status"`
	Volunteer 	string 	`json:"Volunteer" firestore:"volunteer"`
}