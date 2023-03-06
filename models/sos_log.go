package models

type SOSLog struct {
	crId string `firestore:"cr_id,omitempty"`
	datetime int64 `firestore:"datetime,omitempty"`
	startLocation struct{
		lat float64 `firestore:"lat,omitempty"`
		lng float64 `firestore:"lng,omitempty"`
	}
	status string `firestore:"status,omitempty"` //consider declaring enum for this
	volunteerList []string `firestore:"volunteer"`
}