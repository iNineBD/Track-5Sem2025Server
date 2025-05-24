package service

import (
	"inine-track/pkg/database"
	"inine-track/pkg/dto/plataformdto"
	"inine-track/pkg/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetPlatforms() (status int, response gin.H) {

	var platforms []models.DimPlatform
	var listPlataforms []plataformdto.GetPlataformsResponse

	result := database.DB.Find(&platforms)

	if result.Error != nil {
		return http.StatusBadRequest, gin.H{"error": "erro ao buscar plataformas"}
	}

	for _, p := range platforms {
		listPlataforms = append(listPlataforms, plataformdto.GetPlataformsResponse{IDPlatform: p.IDPlatform, NomePlatforme: strings.ToUpper(p.NamePlatform)})
	}

	return http.StatusOK, gin.H{"success": listPlataforms}

}
