package models

type TravelLog struct {
	crId string `firestore:"cr_id,omitempty"`
	datetime int64 `firestore:"datetime,omitempty"`
	currentLocation struct{
		lat float64 `firestore:"lat,omitempty"`
		lng float64 `firestore:"lng,omitempty"`
	}
	status string `firestore:"status,omitempty"` //consider declaring enum for this
}