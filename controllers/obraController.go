package controllers

import (
	"context"
	//"encoding/json"
	"log"
	"net/http"

	"github.com/bgbritodev/MuApp-backend/config"
	"github.com/bgbritodev/MuApp-backend/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var obraCollection *mongo.Collection

func init() {
	mongoClient := config.Connect()
	obraCollection = mongoClient.Database("MuApp").Collection("Obras")
	log.Println("Conectado à collection de obras")
}

// CreateObra cria uma nova obra no banco de dados
func CreateObra(c *gin.Context) {
	var obra models.Obra
	if err := c.BindJSON(&obra); err != nil {
		log.Printf("Error decoding request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding request body"})
		return
	}

	// Atribuir um ID único à obra
	obra.ID = primitive.NewObjectID()

	// Inserir a obra no banco de dados
	_, err := obraCollection.InsertOne(context.TODO(), obra)
	if err != nil {
		log.Printf("Error inserting obra into database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting obra into database"})
		return
	}

	// Responder com sucesso
	c.JSON(http.StatusCreated, obra)
}

// GetObra recupera uma obra específica do banco de dados
func GetObra(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var obra models.Obra
	err = obraCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&obra)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Obra not found"})
		return
	}

	c.JSON(http.StatusOK, obra)
}

// GetObrasBySalaID recupera obras baseadas no ID da sala
func GetObrasBySalaID(c *gin.Context) {
	salaID := c.Param("salaId")

	// Consultar o banco de dados por obras que possuem o salaId correspondente
	filter := bson.M{"salaId": salaID}
	cursor, err := obraCollection.Find(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding obras"})
		return
	}
	defer cursor.Close(context.TODO())

	var obras []models.Obra
	if err = cursor.All(context.TODO(), &obras); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding obras"})
		return
	}

	if len(obras) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No obras found for the given SalaID"})
		return
	}

	c.JSON(http.StatusOK, obras)
}
