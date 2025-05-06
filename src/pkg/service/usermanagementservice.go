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

func UpdateRoleUser(idUser int64, idRole int64) (status int, response gin.H) {

	tx := database.DB.Begin()

	if tx.Error != nil {
		return http.StatusInternalServerError, gin.H{"error": "erro ao iniciar transação"}
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var user models.DimUser

	result := tx.Where("id_user = ?", idUser).First(&user)

	if result.Error != nil {
		tx.Rollback()
		return http.StatusBadRequest, gin.H{"error": "usuário não localizado para atualização"}
	}

	var role models.DimRole

	result = tx.Where("id_role = ?", idRole).First(&role)

	if result.Error != nil {
		tx.Rollback()
		return http.StatusBadRequest, gin.H{"error": "função não localizada para atualização"}
	}

	result = tx.Model(user).Updates(role)

	if result.Error != nil {
		tx.Rollback()
		return http.StatusBadRequest, gin.H{"error": "erro ao atualizar a função do usuário"}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, gin.H{"error": "erro ao completar transação"}
	}

	return http.StatusOK, gin.H{"success": "função atualizada com sucesso"}
}

func GetRoles() (status int, response gin.H) {

	var roles []models.DimRole
	var listRoles []usermanagementdto.GetRoles

	result := database.DB.Find(&roles)

	if result.Error != nil {
		return http.StatusBadRequest, gin.H{"error": "erro ao retornar as roles"}
	}

	for _, role := range roles {
		listRoles = append(listRoles, usermanagementdto.GetRoles{IDRole: role.IDRole, NameRole: strings.ToUpper(role.NameRole)})
	}

	return http.StatusOK, gin.H{"success": listRoles}
}
