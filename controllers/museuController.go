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

var museuCollection *mongo.Collection

func init() {
	mongoClient := config.Connect()
	museuCollection = mongoClient.Database("MuApp").Collection("Museus")
	log.Println("Conectado à coleção de museus")
}

// CreateMuseu cria um novo museu no banco de dados
func CreateMuseu(w http.ResponseWriter, r *http.Request) {
	var museu models.Museu
	err := json.NewDecoder(r.Body).Decode(&museu)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	museu.ID = primitive.NewObjectID()

	_, err = museuCollection.InsertOne(context.TODO(), museu)
	if err != nil {
		http.Error(w, "Error inserting museu into database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(museu)
}

// GetMuseu recupera um museu específico do banco de dados
func GetMuseu(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var museu models.Museu
	err = museuCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&museu)
	if err != nil {
		http.Error(w, "Museu not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(museu)
}

func GetAllMuseus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var museus []models.Museu

	// Buscando todos os museus no MongoDB
	cursor, err := museuCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var museu models.Museu
		if err := cursor.Decode(&museu); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		museus = append(museus, museu)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(museus)
}

// UpdateMuseu atualiza um museu existente no banco de dados
func UpdateMuseu(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var museu models.Museu
	err = json.NewDecoder(r.Body).Decode(&museu)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
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
		http.Error(w, "Error updating museu", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(museu)
}

// DeleteMuseu deleta um museu do banco de dados
func DeleteMuseu(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	_, err = museuCollection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		http.Error(w, "Error deleting museu", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
