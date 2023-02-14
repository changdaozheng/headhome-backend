package models

type MemberStruct struct {
	Name string `firestore:"name,omitempty"`
	Id int64 `firestore:"id,omitempty"`
}