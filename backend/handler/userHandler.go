package handler

import (
	"GoToWeb/backend/model"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type UserHandler struct {
	collection *mongo.Collection
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func NewUserHandler(client *mongo.Client) *UserHandler {
	collection := client.Database("my-database-mongodb").Collection("users")
	return &UserHandler{collection}
}

// CreateUser crée un nouvel utilisateur.
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	// Dans CreateUser
	result, err := h.collection.InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.ID = result.InsertedID.(primitive.ObjectID) // Assurez-vous que le type est correct
	json.NewEncoder(w).Encode(user)
}

// GetUser récupère un utilisateur par ID.
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	log.Println("Recherche de l'utilisateur avec ID:", id)
	var user model.User
	err = h.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// UpdateUser met à jour les données de l'utilisateur.
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)

	// Convertir l'ID de la chaîne en ObjectID
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Mettre à jour l'utilisateur dans la base de données
	filter := bson.M{"_id": id}
	update := bson.M{"$set": user}
	_, err = h.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Récupérer l'utilisateur mis à jour pour le renvoyer
	var updatedUser model.User
	err = h.collection.FindOne(context.Background(), filter).Decode(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedUser)
}
