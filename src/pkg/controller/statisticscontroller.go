package controller

import (
	"inine-track/pkg/dto/statisticsdto"
	"inine-track/pkg/service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Summary Display of all projects registered in taiga
// @Description This endpoint displays all projects based on date range
// @Tags Statistics
// @Param id path string true "Id do projeto para busca"
// @Param data1 query string false "Data de início (formato: YYYY-MM-DD)" example("2025-04-01")
// @Param data2 query string false "Data de fim (formato: YYYY-MM-DD)" example("2025-04-30")
// @Produce json
// @Router /statistics/data/{id} [get]
func GetStatisticsData(c *gin.Context) {

	var idProject, _ = strconv.Atoi(strings.TrimSpace(c.Param("id")))
	var data1 = c.DefaultQuery("data1", "2000-01-01")
	var data2 = c.DefaultQuery("data2", "2025-01-01")

	status, dataTag := service.GetCardsPerTag(idProject, data1, data2)

	if status != 200 {
		c.JSON(status, gin.H{"error": "erro ao retornar as cards por tag"})
		return
	}

	status, dataUser := service.GetCardsPerUser(idProject, data1, data2)

	if status != 200 {
		c.JSON(status, gin.H{"error": "erro ao retornar as cards por colaborador"})
		return
	}

	status, dataStatus := service.GetCardsPerStatus(idProject, data1, data2)

	if status != 200 {
		c.JSON(status, gin.H{"error": "erro ao retornar as cards por tag"})
		return
	}
	status, reworkCards := service.GetTimeDesenvCards(idProject, data1, data2)

	if status != 200 {
		c.JSON(status, gin.H{"error": "erro ao retornar a quantidade de retrabalho por card"})
		return
	}
	var response statisticsdto.GetStatisticsResponse = statisticsdto.GetStatisticsResponse{TagData: dataTag, UserData: dataUser, StatusData: dataStatus, ReworkCards: reworkCards}

	c.JSON(status, gin.H{"success": response})
}
