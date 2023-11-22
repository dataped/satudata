package models

type User struct {
	Name         string `json:"name" bson:"name"`
	Email        string `json:"email" bson:"email"`
	Role         string `json:"role" bson:"role"`
	AccessToken  string `json:"access_token" bson:"access_token"`
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
}

type UserReq struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
