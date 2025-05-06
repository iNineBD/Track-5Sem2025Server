package controller

import (
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
