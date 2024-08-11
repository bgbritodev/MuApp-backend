package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/bgbritodev/MuApp-backend/config"
	"github.com/bgbritodev/MuApp-backend/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var obraCollection *mongo.Collection

func init() {
	mongoTestClient := config.Connect()
	obraCollection = mongoTestClient.Database("MuApp").Collection("Obras")
	log.Println("Connectado a collection")
}

func CreateObras(w http.ResponseWriter, r *http.Request) {
	var obra models.Obra
	err := json.NewDecoder(r.Body).Decode(&obra)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	// Atribuir um ID único à obra
	obra.ID = primitive.NewObjectID()

	// Inserir a obra no banco de dados
	_, err = obraCollection.InsertOne(context.TODO(), obra)
	if err != nil {
		log.Printf("Error inserting obra into database: %v", err)
		http.Error(w, "Error inserting obra into database", http.StatusInternalServerError)
		return
	}

	// Responder com sucesso
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(obra)
}

// GetObra recupera uma obra específica do banco de dados
func GetObra(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var obra models.Obra
	err = obraCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&obra)
	if err != nil {
		http.Error(w, "Obra not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(obra)
}

func GetObrasBySalaID(w http.ResponseWriter, r *http.Request) {
	salaID := mux.Vars(r)["salaId"]

	// Consultar o banco de dados por obras que possuem o salaId correspondente
	filter := bson.M{"salaId": salaID}
	cursor, err := obraCollection.Find(context.TODO(), filter)
	if err != nil {
		http.Error(w, "Error finding obras", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var obras []models.Obra
	if err = cursor.All(context.TODO(), &obras); err != nil {
		http.Error(w, "Error decoding obras", http.StatusInternalServerError)
		return
	}

	if len(obras) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No obras found for the given SalaID")
		return
	}

	json.NewEncoder(w).Encode(obras)
}
