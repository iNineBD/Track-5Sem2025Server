package controller

import (
	"inine-track/pkg/dto/usermanagementdto"
	"inine-track/pkg/service"
	"inine-track/pkg/service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Display of relation user and role
// @Descripition This endpoint displays relation user and role
// @Tags User Management
// @Security BearerAuth
// @Produce json
// @Router /api/usermanagement/data [get]
func GetRelationUserRole(c *gin.Context) {

	_, err := utils.VerifyAndDecodeToken(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	status, response := service.GetRelationUserRole()

	c.JSON(status, response)

}

// @Summary Display of relation all role
// @Descripition This endpoint displays relation all role
// @Tags User Management
// @Security BearerAuth
// @Produce json
// @Router /api/usermanagement/data/roles [get]
func GetRoles(c *gin.Context) {

	_, err := utils.VerifyAndDecodeToken(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	status, response := service.GetRoles()

	c.JSON(status, response)

}

// @Summary Display of relation user and role
// @Description This endpoint displays relation user and role
// @Tags User Management
// @Security BearerAuth
// @Accept json
// @Param usermanagement body usermanagementdto.UpdateRelationUserRole true "Dados para atualizar a role do usuário"
// @Produce json
// @Router /api/usermanagement/data [put]
func UpdateRoleUser(c *gin.Context) {

	_, err := utils.VerifyAndDecodeToken(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	var dataUser usermanagementdto.UpdateRelationUserRole

	if err := c.ShouldBindBodyWithJSON(&dataUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	status, response := service.UpdateRoleUser(dataUser.IDUser, dataUser.IDRole)

	c.JSON(status, response)

}
