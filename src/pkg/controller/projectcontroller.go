package controller

import (
	"inine-track/pkg/service"

	"github.com/gin-gonic/gin"
)

// @Summary Display of all projects registered in taiga
// @Descripition This endpoint displays all projects
// @Tags Projects
// @Produce json
// @Router /projects/data [get]
func GetProjects(c *gin.Context) {

	status, reponse := service.GetProjects()

	c.JSON(status, reponse)
}
