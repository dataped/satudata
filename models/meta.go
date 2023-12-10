package models

type Meta struct {
	Error   int    `json:"error" bson:"error"`
	Message string `json:"message" bson:"message"`
}
