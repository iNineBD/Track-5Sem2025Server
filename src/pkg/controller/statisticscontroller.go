package controller

import (
	"inine-track/pkg/dto/statisticsdto"
	"inine-track/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Display of all projects registered in taiga
// @Descripition This endpoint displays all projects
// @Tags Statistics
// @Produce json
// @Router /statistics/data [get]
func GetStatisticsData(c *gin.Context) {

	var resquest statisticsdto.GetStatisticsRequest

	if err := c.ShouldBindBodyWithJSON(&resquest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	status, dataTag := service.GetCardsPerTag(resquest.Id)

	if status != 200 {
		c.JSON(status, gin.H{"error": "erro ao retornar as cards por tag"})
	}

	status, dataUser := service.GetCardsPerUser(resquest.Id)

	if status != 200 {
		c.JSON(status, gin.H{"error": "erro ao retornar as cards por colaborador"})
	}

	status, dataStatus := service.GetCardsPerStatus(resquest.Id)

	if status != 200 {
		c.JSON(status, gin.H{"error": "erro ao retornar as cards por tag"})
	}

	var response statisticsdto.GetStatisticsResponse = statisticsdto.GetStatisticsResponse{TagData: dataTag, UserData: dataUser, StatusData: dataStatus}

	c.JSON(status, gin.H{"success": response})
}
