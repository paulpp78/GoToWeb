package main

import (
	"GoToWeb/backend/handler"
	"context"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

func main() {

	// Connexion à MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	userHandler := handler.NewUserHandler(client)

	r := mux.NewRouter()
	r.HandleFunc("/user", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/user/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/user", optionsHandler).Methods("OPTIONS")
	r.HandleFunc("/user/{id}", optionsHandler).Methods("OPTIONS")

	log.Println("Le serveur écoute sur le port :3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
func optionsHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	w.WriteHeader(http.StatusOK)
}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
