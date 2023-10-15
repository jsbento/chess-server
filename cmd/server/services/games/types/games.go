package types

import "time"

type Game struct {
	Id       string    `bson:"_id" json:"id"`
	PlayerId string    `bson:"pId" json:"playerId"`
	Result   string    `bson:"res" json:"result"`
	History  []string  `bson:"hist" json:"history"`
	Date     time.Time `bson:"dt" json:"date"`
}

type UpdateGame struct {
	Id       string   `bson:"_id" json:"id"`
	PlayerId *string  `bson:"pId,omitempty" json:"playerId,omitempty"`
	Result   *string  `bson:"result,omitempty" json:"result,omitempty"`
	History  []string `bson:"history" json:"history"`
}
