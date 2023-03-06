package models

type Relationship struct{
	id string `firestore:"id,omitempty"`
	relationship string `firestore:"relationship,omitempt"`
}

type CareGiver struct {
	CgId string `firestore:"cg_id,omitempty"`
	Name string `firestore:"name,omitempty"`
	Address string `firestore:"address,omitempty"`
	ContactNum string `firestore:"contact_num,omitempty"`
	CareReceiver []Relationship  `firestore:"care_receiver,omitempty"`
	ProfilePic string `firestore:"profile_pic,omitempty"`
}