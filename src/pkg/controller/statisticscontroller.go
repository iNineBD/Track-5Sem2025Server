package controller

import (
	"inine-track/pkg/dto/statisticsdto"
	"inine-track/pkg/service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Summary Display of all projects registered in taiga
// @Descripition This endpoint displays all projects
// @Tags Statistics
// @Param id path string true "Id do projeto para busca"
// @Produce json
// @Router /statistics/data/{id} [get]
func GetStatisticsData(c *gin.Context) {

	var idProject, _ = strconv.Atoi(strings.TrimSpace(c.Param("id")))

	status, dataTag := service.GetCardsPerTag(idProject)

	if status != 200 {
		c.JSON(status, gin.H{"error": "erro ao retornar as cards por tag"})
		return
	}

	status, dataUser := service.GetCardsPerUser(idProject)

	if status != 200 {
		c.JSON(status, gin.H{"error": "erro ao retornar as cards por colaborador"})
		return
	}

	status, dataStatus := service.GetCardsPerStatus(idProject)

	if status != 200 {
		c.JSON(status, gin.H{"error": "erro ao retornar as cards por tag"})
		return
	}

	var response statisticsdto.GetStatisticsResponse = statisticsdto.GetStatisticsResponse{TagData: dataTag, UserData: dataUser, StatusData: dataStatus}

	c.JSON(status, gin.H{"success": response})
}
