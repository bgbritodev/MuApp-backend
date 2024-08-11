package controllers

import (
	"context"
	"encoding/json"
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
	obraCollection = config.Client.Database("MuApp").Collection("Obras")
}

// CreateObra cria uma nova obra no banco de dados
func CreateObra(w http.ResponseWriter, r *http.Request) {
	var obra models.Obra
	err := json.NewDecoder(r.Body).Decode(&obra)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	obra.ID = primitive.NewObjectID()
	_, err = obraCollection.InsertOne(context.TODO(), obra)
	if err != nil {
		http.Error(w, "Error inserting obra into database", http.StatusInternalServerError)
		return
	}

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
