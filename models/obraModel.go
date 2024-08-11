package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Obra struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Autor       string             `bson:"autor"`
	Description string             `bson:"description"`
	Image       string             `bson:"image"`
	Audio       string             `bson:"audio"`
}
