package models

type CareReciever struct {
	CrId string `firestore:"cr_id,omitempty"`
	Name string `firestore:"name,omitempty"`
	Address string `firestore:"address,omitempty"`
	ContactNum string `firestore:"contact_num,omitempty"`
	SafezoneCtr struct {
		Lat float64 `firestore:"lat,omitempty"`
		Lng float64 `firestoer:"lng,omitempty"`
	}
	SafezoneRadius int64 `firestore:"safezone_radius,omitempty"`
	CareGiver []Relationship  `firestore:"care_giver,omitempty"` //Relationship is a self-declared struct in care_giver.go
	ProfilePic string `firestore:"profile_pic,omitempty"`
}