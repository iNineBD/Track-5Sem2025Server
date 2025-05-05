package controller

import (
	"inine-track/pkg/service"
	"inine-track/pkg/service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Display of all projects registered in taiga
// @Descripition This endpoint displays all projects
// @Tags Projects
// @Security BearerAuth
// @Produce json
// @Router /api/projects/data [get]
func GetProjects(c *gin.Context) {

	claims, err := utils.VerifyAndDecodeToken(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	idUser := int64(claims["user_id"].(float64))
	idRole := int64(claims["role"].(float64))

	status, reponse := service.GetProjects(idUser, idRole)

	c.JSON(status, reponse)
}
