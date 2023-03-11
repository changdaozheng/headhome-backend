package models

type Relationship struct{
	Id string `firestore:"id"`
	Relationship string `firestore:"relationship"`
}

type CareGiver struct {
	CgId string `firestore:"cg_id"`
	Name string `firestore:"name"`
	Address string `firestore:"address"`
	ContactNum string `firestore:"contact_num"`
	CareReceiver []Relationship  `firestore:"care_receiver"`
	ProfilePic string `firestore:"profile_pic"`
}