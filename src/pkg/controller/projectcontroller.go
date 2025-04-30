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

	_, err := utils.VerifyAndDecodeToken(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	status, reponse := service.GetProjects()

	c.JSON(status, reponse)
}
