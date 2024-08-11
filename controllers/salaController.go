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

var salaCollection *mongo.Collection

func init() {
	mongoTestClient := config.Connect()
	salaCollection = mongoTestClient.Database("MuApp").Collection("Salas")
	log.Println("Conectado à coleção de salas")
}

// CreateSala cria uma nova sala no banco de dados
func CreateSala(w http.ResponseWriter, r *http.Request) {
	var sala models.Sala
	err := json.NewDecoder(r.Body).Decode(&sala)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	// Atribuir um ID único à sala
	sala.ID = primitive.NewObjectID()

	// Inserir a sala no banco de dados
	_, err = salaCollection.InsertOne(context.TODO(), sala)
	if err != nil {
		log.Printf("Error inserting sala into database: %v", err)
		http.Error(w, "Error inserting sala into database", http.StatusInternalServerError)
		return
	}

	// Responder com sucesso
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sala)
}

// GetSalasByMuseuID recupera todas as salas de um determinado museu
func GetSalasByMuseuID(w http.ResponseWriter, r *http.Request) {
	museuID := mux.Vars(r)["museuId"]

	// Consultar o banco de dados por salas que possuem o museuId correspondente
	filter := bson.M{"museuId": museuID}
	cursor, err := salaCollection.Find(context.TODO(), filter)
	if err != nil {
		http.Error(w, "Error finding salas", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var salas []models.Sala
	if err = cursor.All(context.TODO(), &salas); err != nil {
		http.Error(w, "Error decoding salas", http.StatusInternalServerError)
		return
	}

	if len(salas) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No salas found for the given MuseuID")
		return
	}

	json.NewEncoder(w).Encode(salas)
}

func GetSala(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var sala models.Sala
	err = salaCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&sala)
	if err != nil {
		http.Error(w, "sala not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(sala)
}
