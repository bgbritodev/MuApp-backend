package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sala struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Order       int                `bson:"order"`
	Image       string             `bson:"image"`
	MuseuID     string             `bson:"museuId"`
}
