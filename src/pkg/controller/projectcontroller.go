package controller

import (
	"inine-track/pkg/service"
	"inine-track/pkg/service/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Display of all projects registered in taiga
// @Description This endpoint displays all projects
// @Tags Projects
// @Security BearerAuth
// @Param idPlatform path string true "Id da plataforma para busca"
// @Produce json
// @Router /api/projects/data/{idPlatform} [get]
func GetProjects(c *gin.Context) {

	claims, err := utils.VerifyAndDecodeToken(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	var idPlatform, _ = strconv.ParseInt(c.Param("idPlatform"), 10, 64)
	idUser := int64(claims["user_id"].(float64))
	idRole := int64(claims["role"].(float64))

	status, reponse := service.GetProjects(idUser, idRole, idPlatform)

	c.JSON(status, reponse)
}
