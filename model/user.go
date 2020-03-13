package model

type User struct {
	UserName string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password,omitempty"`
}

type JwtToken struct {
	Token string `json:"token"`
}
