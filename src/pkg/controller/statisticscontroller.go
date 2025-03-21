package controller

import (
	"inine-track/pkg/dto/statisticsdto"
	"inine-track/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStatisticsData(c *gin.Context) {

	var resquest statisticsdto.GetStatisticsRequest

	if err := c.ShouldBindBodyWithJSON(resquest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	status, response := service.GetStatisticsData(resquest.Id)

	c.JSON(status, response)
}
