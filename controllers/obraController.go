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

// CreateObras cria múltiplas obras no banco de dados
func CreateObras(w http.ResponseWriter, r *http.Request) {
	var obras []models.Obra
	err := json.NewDecoder(r.Body).Decode(&obras)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	for i := range obras {
		obras[i].ID = primitive.NewObjectID()
	}

	_, err = obraCollection.InsertMany(context.TODO(), toInterfaceSlice(obras))
	if err != nil {
		http.Error(w, "Error inserting obras into database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(obras)
}

// Função auxiliar para converter um slice de structs para um slice de interface{}
func toInterfaceSlice(obras []models.Obra) []interface{} {
	var result []interface{}
	for _, obra := range obras {
		result = append(result, obra)
	}
	return result
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
