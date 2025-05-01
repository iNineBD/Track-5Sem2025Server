package service

import (
	"inine-track/pkg/database"
	"inine-track/pkg/dto/projectdto"
	"inine-track/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProjects(idUser int64) (int, gin.H) {
	var listProjects []projectdto.GetProjectsResponse
	var projects []models.DimProject

	result := database.DB.Raw(`SELECT dc.id_card, dc.name_card FROM fato_cards fc
    INNER JOIN dim_card dc ON dc.id_card = fc.id_card WHERE fc.id_user = $1
    GROUP BY dc.id_card`, idUser).Find(&projects)

	if result.Error != nil {
		return http.StatusBadRequest, gin.H{"error": "erro ao buscar projetos"}
	}

	if len(projects) == 0 {
		return http.StatusNotFound, gin.H{"error": "não existem projetos cadastrados"}
	}

	for _, p := range projects {
		var project projectdto.GetProjectsResponse = projectdto.GetProjectsResponse{ID: p.IDProject, Name: p.NameProject, Description: p.Description}
		listProjects = append(listProjects, project)
	}

	return http.StatusOK, gin.H{"success": listProjects}
}
