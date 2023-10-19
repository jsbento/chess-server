package types

import "time"

type User struct {
	Id        string    `bson:"_id" json:"id"`
	Username  string    `bson:"username" json:"username"`
	Email     string    `bson:"email" json:"email"`
	Password  string    `bson:"password" json:"-"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}

type UpdateUser struct {
	Id       string  `bson:"_id" json:"id"`
	Username *string `bson:"username,omitempty" json:"username,omitempty"`
	Email    *string `bson:"email,omitempty" json:"email,omitempty"`
}
