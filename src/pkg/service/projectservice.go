package service

import (
	"inine-track/pkg/database"
	"inine-track/pkg/dto/projectdto"
	"inine-track/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProjects() (int, gin.H) {
	var listProjects []projectdto.GetProjectsResponse
	var projects []models.DimProject

	result := database.DB.Find(&projects)

	if result.Error != nil {
		return http.StatusBadRequest, gin.H{"error": "erro ao buscar projetos"}
	}

	if len(projects) == 0 {
		return http.StatusNotFound, gin.H{"error": "não existem projetos cadastrados"}
	}

	for _, p := range projects {
		var project projectdto.GetProjectsResponse = projectdto.GetProjectsResponse{Name: p.Name, Description: p.Description,
			CreatedDate: p.CreatedDate, ModifiedDate: p.ModifiedDate, FinishDate: p.FinishDate, LogoBigUrl: p.LogoBigUrl, LogoSmallUrl: p.LogoSmallUrl}
		listProjects = append(listProjects, project)
	}

	return http.StatusOK, gin.H{"success": listProjects}
}
