package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Media representa a estrutura de um documento de mídia no MongoDB
type Media struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Type string             `bson:"type" json:"type"`
	Name string             `bson:"name" json:"name"`
	URL  string             `bson:"url" json:"url"`
}
