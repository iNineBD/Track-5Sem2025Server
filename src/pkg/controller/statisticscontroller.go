package controller

import (
	"inine-track/pkg/service"
	"inine-track/pkg/service/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Display of all projects registered in taiga
// @Description This endpoint displays all projects based on date range
// @Tags Statistics
// @Security BearerAuth
// @Param id path string true "Id do projeto para busca"
// @Param data1 query string false "Data de início (formato: YYYY-MM-DD)" example("2025-04-01")
// @Param data2 query string false "Data de fim (formato: YYYY-MM-DD)" example("2025-04-30")
// @Produce json
// @Router /api/statistics/data/{id} [get]
func GetStatisticsData(c *gin.Context) {

	claims, err := utils.VerifyAndDecodeToken(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	idRole := int64(claims["role"].(float64))
	idUser := int64(claims["user_id"].(float64))

	var idProject, _ = strconv.ParseInt(c.Param("id"), 10, 64)
	var data1 = c.DefaultQuery("data1", "2000-01-01")
	var data2 = c.DefaultQuery("data2", "2025-01-01")

	status, response := service.GetMetrics(idProject, data1, data2, idUser, idRole)

	c.JSON(int(status), response)
}
