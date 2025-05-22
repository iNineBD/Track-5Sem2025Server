package service

import (
	"fmt"
	"inine-track/pkg/database"
	"inine-track/pkg/dto/projectdto"
	"inine-track/pkg/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetProjects(idUser int64, idRole int64, idPlatform int64) (int, gin.H) {
	var listProjects []projectdto.GetProjectsResponse
	var projects []models.DimProject
	var role models.DimRole

	result := database.DB.Where("id_role = ?", idRole).First(&role)

	if result.Error != nil {
		return http.StatusBadRequest, gin.H{"error": "erro ao trazer a role do usuário"}
	}
	var roleName string = role.NameRole

	if strings.ToUpper(roleName) == "ADMIN" {
		result = database.DB.Preload("DimPlatform").Where("id_platform = ?", idPlatform).Find(&projects)
	} else {
		result = database.DB.Raw(`SELECT dp.id_project, dp.name_project, dp.description FROM fato_cards fc
		INNER JOIN dim_project dp ON fc.id_project = dp.id_project
		INNER JOIN dim_platform dpa on dpa.id_platform = dp.id_platform WHERE fc.id_user = $1 and dp.id_platform = $2
		GROUP BY dp.id_project`, idUser, idPlatform).Find(&projects)
	}

	fmt.Println(projects)

	if result.Error != nil {
		return http.StatusBadRequest, gin.H{"error": "erro ao buscar projetos"}
	}

	if len(projects) == 0 {
		return http.StatusNotFound, gin.H{"error": "não existem projetos cadastrados"}
	}

	for _, p := range projects {
		if p.IDProject != 0 {
			var project projectdto.GetProjectsResponse = projectdto.GetProjectsResponse{ID: p.IDProject, Name: p.NameProject, Description: p.Description}
			listProjects = append(listProjects, project)
		}
	}

	return http.StatusOK, gin.H{"success": listProjects}
}
