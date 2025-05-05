package repository

import (
	"context"
	"fmt"
	"go_portfolio/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// CitationUpdateResponse definie the response structure
type CitationUpdateResponse struct {
	message string
}

// CitationResponse define the response structure
type CitationResponse struct {
	Data     []bson.M `json:"data"`
	Metadata struct {
		Total    int `json:"total"`
		Page     int `json:"page"`
		PageSize int `json:"page_size"`
	} `json:"metadata"`
}

// CitationRepository définit l'interface pour les opérations sur les citations
type CitationRepository interface {
	GetTheCitation(ctx context.Context, page, pageSize int) (*CitationResponse, error)
	UpdateCitation(id string, citation *models.Citation) (*CitationUpdateResponse, error)
}

// citationRepository implémente l'interface CitationRepository
type citationRepository struct {
	db *mongo.Database
}

// NewCitationRepository crée une nouvelle instance de citationRepository
func NewCitationRepository(db *mongo.Database) CitationRepository {
	return &citationRepository{db: db}
}

// GetTheCitations implémente la méthode pour récupérer les citations
func (r *citationRepository) GetTheCitation(ctx context.Context, page, pageSize int) (*CitationResponse, error) {
	// Logging de la requête
	log.Printf("Requête GET /citations reçue")

	// Vérification de la connexion
	if r.db == nil {
		return nil, fmt.Errorf("base de données non disponible")
	}

	// Recherche dans la collection
	collection := r.db.Collection("top_citation")
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
	response := &CitationResponse{
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

// UpdateCitation implémente la méthode pour mettre à jour une citation
func (r *citationRepository) UpdateCitation(id string, citation *models.Citation) (*CitationUpdateResponse, error) {
	// Vérification de la connexion
	if r.db == nil {
		return nil, fmt.Errorf("base de données non disponible")
	}

	// Conversion de l'ID en string en ObjectID
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("ID invalide: %v", err)
	}

	// Vérification de la collection
	collection := r.db.Collection("top_citation")
	if collection == nil {
		return nil, fmt.Errorf("collection non trouvée")
	}

	// Création de l'opération de mise à jour
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "citation", Value: citation.Citation},
			{Key: "author", Value: citation.Author},
			{Key: "updated_at", Value: time.Now()},
		}},
	}

	// Exécution de la mise à jour
	result, err := collection.UpdateOne(
		context.TODO(),
		bson.D{{Key: "_id", Value: objID}},
		update,
	)

	if err != nil {
		return nil, fmt.Errorf("erreur lors de la mise à jour: %v", err)
	}

	// Vérification que la mise à jour a été effectuée
	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("citation non trouvée ou non modifiée")
	}

	response := &CitationUpdateResponse{
		message: "success",
	}

	return response, nil
}
