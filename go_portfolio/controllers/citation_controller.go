package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"go_portfolio/databases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CitationResponse struct {
	Data     []bson.M `json:"data"`
	Metadata struct {
		Total    int `json:"total"`
		Page     int `json:"page"`
		PageSize int `json:"page_size"`
	} `json:"metadata"`
}

func GetAllCitations() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logging de la requête
		log.Printf("Requête GET /citations reçue")

		// Context avec timeout
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Vérification de la connexion
		if databases.Database == nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Base de données non disponible",
			})
			return
		}

		// Recherche dans la collection
		collection := databases.Database.Collection("top_citation")
		if collection == nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Collection non trouvée",
			})
			return
		}

		// Configuration de la pagination
		page, pageSize := 1, 10
		if p := c.Query("page"); p != "" {
			page = parseInt(p)
		}
		if ps := c.Query("page_size"); ps != "" {
			pageSize = parseInt(ps)
		}

		// Création du filtre et de l'option de pagination
		opts := options.Find()
		opts.SetSkip(int64((page - 1) * pageSize))
		opts.SetLimit(int64(pageSize))

		// Exécution de la requête
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Erreur lors de la recherche: %v", err),
			})
			return
		}
		defer cursor.Close(ctx)

		// Récupération des résultats
		var results []bson.M
		if err = cursor.All(ctx, &results); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Erreur lors du décodage: %v", err),
			})
			return
		}

		// Création de la réponse structurée
		response := CitationResponse{
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

		c.JSON(http.StatusOK, response)
	}
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
