package repository

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// AboutMeResponse define the response structure
type AboutMeResponse struct {
	Data     []bson.M `json:"data"`
	Metadata struct {
		Total    int `json:"total"`
		Page     int `json:"page"`
		PageSize int `json:"page_size"`
	} `json:"metadata"`
}

// AboutMeRepository définit l'interface pour les opérations sur les AboutMes
type AboutMeRepository interface {
	GetTheAboutMe(ctx context.Context, page, pageSize int) (*AboutMeResponse, error)
}

// AboutMeRepository implémente l'interface AboutMeRepository
type aboutmeRepository struct {
	db *mongo.Database
}

// NewAboutMeRepository crée une nouvelle instance de AboutMeRepository
func NewAboutMeRepository(db *mongo.Database) AboutMeRepository {
	return &aboutmeRepository{db: db}
}

// GetTheAboutMes implémente la méthode pour récupérer les AboutMes
func (r *aboutmeRepository) GetTheAboutMe(ctx context.Context, page, pageSize int) (*AboutMeResponse, error) {
	// Logging de la requête
	log.Printf("Requête GET /AboutMes reçue")

	// Vérification de la connexion
	if r.db == nil {
		return nil, fmt.Errorf("base de données non disponible")
	}

	// Recherche dans la collection
	collection := r.db.Collection("about_me")
	if collection == nil {
		return nil, fmt.Errorf("collection non trouvée")
	}

	// Configuration de la pagination
	opts := options.Find()
	opts.SetSkip(int64((page - 1) * pageSize))
	opts.SetLimit(int64(pageSize))

	// Exécution de la requête
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la recherche: %v", err)
	}
	defer cursor.Close(ctx)

	// Récupération des résultats
	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("erreur lors du décodage: %v", err)
	}

	// Création de la réponse structurée
	response := &AboutMeResponse{
		Data: results,
		Metadata: struct {
			Total    int `json:"total"`
			Page     int `json:"page"`
			PageSize int `json:"page_size"`
		}{
			Total:    len(results),
			Page:     page,
			PageSize: pageSize,
		},
	}

	return response, nil
}
