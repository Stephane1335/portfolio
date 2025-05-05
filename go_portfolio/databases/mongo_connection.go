package databases

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	client   *mongo.Client
	Database *mongo.Database
	ctx      context.Context
)

func ConnectDatabase() error {
	// Création du context avec timeout
	log.Printf("Elle se lance !")

	// Configuration de la connexion
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://yapojeanstephane:5n3cbcWSZvzpXYbv@portfolio.pzp1qkw.mongodb.net/?retryWrites=true&w=majority&appName=portfolio").
		SetServerAPIOptions(serverAPI).
		SetRetryWrites(true).
		SetRetryReads(true)

	// Connexion au serveur
	var err error
	client, err = mongo.Connect(opts)
	if err != nil {
		return fmt.Errorf("erreur de connexion à MongoDB: %w", err)
	}

	// Initialisation de la base de données
	Database = client.Database("portfolio")
	if Database == nil {
		return fmt.Errorf("base de données non initialisée")
	}

	log.Println("Connexion MongoDB établie avec succès")
	log.Println(Database.Collection("top_citation"))
	return nil
}

func DisconnectDatabase() {
	if client != nil {
		if err := client.Disconnect(ctx); err != nil {
			log.Printf("Erreur lors de la déconnexion de MongoDB: %v", err)
		}
	}
}
