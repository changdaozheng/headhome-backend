package models

type CareReceiver struct {
	CrId 			string 	`firestore:"cr_id"`
	Name 			string 	`firestore:"name"`
	Address 		string 	`firestore:"address"`
	ContactNum 		string 	`firestore:"contact_num"`
	Notes 			string 	`firestore:"notes"`
	SafezoneCtr struct {
		Lat float64 `firestore:"lat"`
		Lng float64 `firestore:"lng"`
	} `firestore:"safezone_ctr"`
	SafezoneRadius 	int64 	`firestore:"safezone_radius"`
	CareGiver 		[]Relationship  `firestore:"care_giver"` //Relationship is a self-declared struct in care_giver.go
	ProfilePic 		string 	`firestore:"profile_pic"`
	AuthID 			string	`firestore:"auth_id"`
}