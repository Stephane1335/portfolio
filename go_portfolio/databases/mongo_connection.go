package databases

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var (
	client     *mongo.Client
	Database   *mongo.Database
	ctx        context.Context
	cancelFunc context.CancelFunc
)

func ConnectDatabase() error {
	// Création du context avec timeout
	ctx, cancelFunc = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

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

	// Vérification de la connexion
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return fmt.Errorf("erreur de ping MongoDB: %w", err)
	}

	// Initialisation de la base de données
	Database = client.Database("portfolio")
	if Database == nil {
		return fmt.Errorf("base de données non initialisée")
	}

	log.Println("Connexion MongoDB établie avec succès")
	return nil
}

func DisconnectDatabase() {
	if client != nil {
		if err := client.Disconnect(ctx); err != nil {
			log.Printf("Erreur lors de la déconnexion de MongoDB: %v", err)
		}
	}
}
