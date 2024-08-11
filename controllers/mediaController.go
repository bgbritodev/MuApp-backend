package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/bgbritodev/MuApp-backend/config"
	"github.com/bgbritodev/MuApp-backend/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var mediaCollection *mongo.Collection

func init() {
	mediaCollection = config.Client.Database("MuApp").Collection("media")
}

// UploadMedia lida com o upload de mídia e salva no MongoDB
func UploadMedia(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error getting file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Cria um arquivo temporário para salvar o upload
	tempFile, err := os.CreateTemp("", "upload-*.tmp")
	if err != nil {
		http.Error(w, "Error creating temp file", http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	// Copia o conteúdo do arquivo enviado para o arquivo temporário
	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	// Salvar no MongoDB
	media := models.Media{
		ID:   primitive.NewObjectID(),
		Name: filepath.Base(tempFile.Name()),
		URL:  fmt.Sprintf("/media/%s", filepath.Base(tempFile.Name())),
	}

	_, err = mediaCollection.InsertOne(context.TODO(), media)
	if err != nil {
		http.Error(w, "Error saving media", http.StatusInternalServerError)
		return
	}

	// Retornar resposta de sucesso
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(media)
}

// GetMedia recupera uma mídia específica do MongoDB
func GetMedia(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var media models.Media
	err = mediaCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&media)
	if err != nil {
		http.Error(w, "Media not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, media.URL)
}
