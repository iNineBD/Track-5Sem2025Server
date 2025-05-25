package controller

import (
	"inine-track/pkg/service"
	"inine-track/pkg/service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Display of all platforms registered
// @Description This endpoint displays all platforms
// @Tags Plataforms
// @Security BearerAuth
// @Produce json
// @Router /api/platforms/data [get]
func GetPlatforms(c *gin.Context) {

	_, err := utils.VerifyAndDecodeToken(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	status, reponse := service.GetPlatforms()

	c.JSON(status, reponse)
}
