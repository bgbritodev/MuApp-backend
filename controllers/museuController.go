package controllers

import (
	"context"
	//"encoding/json"
	"log"

	"github.com/bgbritodev/MuApp-backend/config"
	"github.com/bgbritodev/MuApp-backend/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var museuCollection *mongo.Collection

func init() {
	mongoClient := config.Connect()
	museuCollection = mongoClient.Database("MuApp").Collection("Museus")
	log.Println("Conectado à coleção de museus")
}

// CreateMuseu cria um novo museu no banco de dados
func CreateMuseu(c *gin.Context) {
	var museu models.Museu
	if err := c.ShouldBindJSON(&museu); err != nil {
		c.JSON(400, gin.H{"error": "Error decoding request body"})
		return
	}

	museu.ID = primitive.NewObjectID()

	_, err := museuCollection.InsertOne(context.TODO(), museu)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error inserting museu into database"})
		return
	}

	c.JSON(201, museu)
}

// GetMuseu recupera um museu específico do banco de dados
func GetMuseu(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	var museu models.Museu
	err = museuCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&museu)
	if err != nil {
		c.JSON(404, gin.H{"error": "Museu not found"})
		return
	}

	c.JSON(200, museu)
}

// GetAllMuseus recupera todos os museus do banco de dados
func GetAllMuseus(c *gin.Context) {
	var museus []models.Museu

	cursor, err := museuCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var museu models.Museu
		if err := cursor.Decode(&museu); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		museus = append(museus, museu)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, museus)
}

// UpdateMuseu atualiza um museu existente no banco de dados
func UpdateMuseu(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	var museu models.Museu
	if err := c.ShouldBindJSON(&museu); err != nil {
		c.JSON(400, gin.H{"error": "Error decoding request body"})
		return
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"nome":        museu.Name,
			"localizacao": museu.Location,
			"descricao":   museu.Description,
			"image":       museu.Image,
		},
	}

	_, err = museuCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error updating museu"})
		return
	}

	c.JSON(200, museu)
}

// DeleteMuseu deleta um museu do banco de dados
func DeleteMuseu(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	_, err = museuCollection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		c.JSON(500, gin.H{"error": "Error deleting museu"})
		return
	}

	c.Status(204)
}
