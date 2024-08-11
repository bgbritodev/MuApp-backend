package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Museu struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Location    string             `bson:"location"`
	Description string             `bson:"description"`
	Image       string             `bson:"image"`
}
