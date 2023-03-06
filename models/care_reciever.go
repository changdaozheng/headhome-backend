package models

type CareReciever struct {
	crId string `firestore:"cr_id,omitempty"`
	name string `firestore:"name,omitempty"`
	address string `firestore:"address,omitempty"`
	contactNum string `firestore:"contact_num,omitempty"`
	safezoneCtr struct {
		lat float64 `firestore:"lat,omitempty"`
		lng float64 `firestoer:"lng,omitempty"`
	}
	safezoneRadius int64 `firestore:"safezone_radius,omitempty"`
	careGiver []Relationship  `firestore:"care_giver,omitempty"` //Relationship is a self-declared struct in care_giver.go
	profilePic string `firestore:"profile_pic,omitempty"`
}