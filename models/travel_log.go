package models

type TravelLog struct {
	CrId string `firestore:"cr_id"`
	Datetime int64 `firestore:"datetime"`
	CurrentLocation struct{
		Lat float64 `firestore:"lat"`
		Lng float64 `firestore:"lng"`
	}
	Status string `firestore:"status"` //consider declaring enum for this
}