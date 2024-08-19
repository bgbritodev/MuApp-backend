package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/bgbritodev/MuApp-backend/config"
	"github.com/bgbritodev/MuApp-backend/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var salaCollection *mongo.Collection

func init() {
	mongoTestClient := config.Connect()
	salaCollection = mongoTestClient.Database("MuApp").Collection("Salas")
	log.Println("Conectado à coleção de salas")
}

// CreateSala cria uma nova sala no banco de dados
func CreateSala(c *gin.Context) {
	var sala models.Sala
	if err := c.BindJSON(&sala); err != nil {
		log.Printf("Error decoding request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding request body"})
		return
	}

	// Atribuir um ID único à sala
	sala.ID = primitive.NewObjectID()

	// Inserir a sala no banco de dados
	_, err := salaCollection.InsertOne(context.TODO(), sala)
	if err != nil {
		log.Printf("Error inserting sala into database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting sala into database"})
		return
	}

	// Responder com sucesso
	c.JSON(http.StatusCreated, sala)
}

// GetSalasByMuseuID recupera todas as salas de um determinado museu
func GetSalasByMuseuID(c *gin.Context) {
	museuID := c.Param("museuId")

	// Consultar o banco de dados por salas que possuem o museuId correspondente
	filter := bson.M{"museuId": museuID}
	cursor, err := salaCollection.Find(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding salas"})
		return
	}
	defer cursor.Close(context.TODO())

	var salas []models.Sala
	if err = cursor.All(context.TODO(), &salas); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding salas"})
		return
	}

	if len(salas) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No salas found for the given MuseuID"})
		return
	}

	c.JSON(http.StatusOK, salas)
}

// GetSala recupera uma sala específica do banco de dados
func GetSala(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var sala models.Sala
	err = salaCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&sala)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sala not found"})
		return
	}

	c.JSON(http.StatusOK, sala)
}
