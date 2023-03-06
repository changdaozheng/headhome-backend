package models

type Relationship struct{
	id string `firestore:"id,omitempty"`
	relationship string `firestore:"relationship,omitempt"`
}

type CareGiver struct {
	cgId string `firestore:"cg_id,omitempty"`
	name string `firestore:"name,omitempty"`
	address string `firestore:"address,omitempty"`
	contactNum string `firestore:"contact_num,omitempty"`
	careReceiver []Relationship  `firestore:"care_receiver,omitempty"`
	profilePic string `firestore:"profile_pic,omitempty"`
}