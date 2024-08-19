package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/bgbritodev/MuApp-backend/config"
	"github.com/bgbritodev/MuApp-backend/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection

func init() {
	mongoTestClient := config.Connect()
	userCollection = mongoTestClient.Database("MuApp").Collection("Users")
	log.Println("Conectado à coleção de usuários")
}

// Chave secreta para assinar o token JWT (mantenha isso em um lugar seguro)
var jwtKey = []byte("your_secret_key")

// Claims é uma struct que será codificada em um token JWT
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// Função para gerar o hash da senha
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Função para comparar a senha inserida com o hash armazenado
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// CreateUser cria um novo usuário
func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Gera o hash da senha
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar senha"})
		return
	}
	user.Password = hashedPassword
	user.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

// UpdateUser atualiza um usuário existente
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Se a senha for fornecida, hash a nova senha
	if user.Password != "" {
		hashedPassword, err := hashPassword(user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar senha"})
			return
		}
		user.Password = hashedPassword
	} else {
		// Se a senha não for fornecida, não altere o campo password
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var existingUser models.User
		err := userCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&existingUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao encontrar usuário"})
			return
		}
		user.Password = existingUser.Password
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": user,
	}

	_, err := userCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar usuário"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário atualizado com sucesso"})
}

// LoginUser permite que um usuário faça login
func LoginUser(c *gin.Context) {
	var credentials models.User
	var foundUser models.User

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := userCollection.FindOne(ctx, bson.M{"email": credentials.Email}).Decode(&foundUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email ou senha incorretos"})
		return
	}

	// Verifica se a senha está correta
	if !checkPasswordHash(credentials.Password, foundUser.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email ou senha incorretos"})
		return
	}

	// Gerando um token JWT válido por 24 horas
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: credentials.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
