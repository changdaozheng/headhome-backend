package models

type Volunteer struct {
	VId string `firestore:"v_id"`
	Name string `firestore:"name"`
	ContactNum string `firestore:"contact_num"`
	CertificationStart int64 `firestore:"certification_start"`
	CertificationEnd int64 `firestore:"certification_end"`
	ProfilePic string `firestore:"profile_pic"`
}