package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name,omitempty"`
	LastName string             `bson:"lastname" json:"lastname,omitempty"`
	BornDate time.Time          `bson:"borndate" json:"borndate,omitempty"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password,omitempty"`
	Avatar   string             `bson:"avatar" json:"avatar,omitempty"`
	Banner   string             `bson:"banner" json:"banner,omitempty"`
	Biograph string             `bson:"biograph" json:"biograph,omitempty"`
	Location string             `bson:"location" json:"location,omitempty"`
	WebSite  string             `bson:"webtsite" json:"website,omitempty"`
}
