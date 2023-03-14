package models

type TravelLog struct {
	CrId string `json:"CrId" firestore:"cr_id"`
	Datetime int64 `json:"Datetime" firestore:"datetime"`
	CurrentLocation struct{
		Lat float64 `json:"Lat" firestore:"lat"`
		Lng float64 `json:"Lng" firestore:"lng"`
	}
	Status string `json:"Status" firestore:"status"` //consider declaring enum for this
}