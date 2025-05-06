package service

import (
	"inine-track/pkg/database"
	"inine-track/pkg/dto/usermanagementdto"
	"inine-track/pkg/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetRelationUserRole() (status int, response gin.H) {

	var users []models.DimUser
	var listRelationUserRole []usermanagementdto.RelationUserRoleResponse

	result := database.DB.Preload("DimRole").Find(&users)

	if result.Error != nil {
		return http.StatusBadRequest, gin.H{"error": "erro ao retornar a relação usuários e suas funções"}
	}

	for _, user := range users {
		var userRole usermanagementdto.RelationUserRoleResponse = usermanagementdto.RelationUserRoleResponse{IDUser: user.IDUser, NameUser: strings.ToUpper(user.NameUser),
			IDRole: user.IDRole, NameRole: strings.ToUpper(user.DimRole.NameRole)}
		listRelationUserRole = append(listRelationUserRole, userRole)
	}

	return http.StatusOK, gin.H{"success": listRelationUserRole}
}
