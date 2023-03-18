package models

type TravelLog struct {
	CrId		string 	`json:"CrId" firestore:"cr_id"`
	Datetime	int64 	`json:"Datetime" firestore:"datetime"`
	TravelLogId	string	`json:"TravelLogId" firestore:"travel_log_id"`
	CurrentLocation struct{
		Lat float64 `json:"Lat" firestore:"lat"`
		Lng float64 `json:"Lng" firestore:"lng"`
	} `json:"CurrentLocation" firestore:"current_location"`
	Status 		string	`json:"Status" firestore:"status"` //consider declaring enum for this
}