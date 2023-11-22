package models

type ResponseUser struct {
	Meta Meta `json:"meta" bson:"meta"`
	Data User `json:"data" bson:"data"`
}
