package models

type SOSLog struct {
	CrId string `firestore:"cr_id,omitempty"`
	Datetime int64 `firestore:"datetime,omitempty"`
	StartLocation struct{
		Lat float64 `firestore:"lat,omitempty"`
		Lng float64 `firestore:"lng,omitempty"`
	}
	Status string `firestore:"status,omitempty"` //consider declaring enum for this
	VolunteerList []string `firestore:"volunteer"`
}