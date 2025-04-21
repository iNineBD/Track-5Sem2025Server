package service

import (
	"fmt"
	"inine-track/pkg/database"
	"inine-track/pkg/dto/statisticsdto"
	"inine-track/pkg/service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMetrics(IDProject int, data1 string, data2 string) (status int, response gin.H) {

	var listCardsPerTag []statisticsdto.TagData
	var listCardsPerUser []statisticsdto.UserData
	var listCardsPerStatus []statisticsdto.StatusData
	var listCardsRework []statisticsdto.ReworkCards
	var listCardsFinished []statisticsdto.FinishedCards
	var listCardsTimeExecution []statisticsdto.TimeExecutionCards

	err := utils.GetProject(int64(IDProject))

	if err != nil {
		return http.StatusBadRequest, gin.H{"error": err.Error()}
	}

	t1, t2, err := utils.FormateDate(data1, data2)

	if err != nil {
		return http.StatusBadRequest, gin.H{"error": err}
	}

	result := database.DB.Raw(`select * from get_qtd_cards_por_tag($1,$2,$3)`, IDProject, t1, t2).Find(&listCardsPerTag)

	if result.Error != nil {
		return http.StatusBadRequest, gin.H{"error": "erro ao retornar as cards por tag"}
	}

	result = database.DB.Raw(`select * from get_qtd_cards_por_colaborador($1,$2,$3)`, IDProject, t1, t2).Find(&listCardsPerUser)

	if result.Error != nil {
		return http.StatusBadRequest, gin.H{"error": "erro ao retornar as cards por colaborador"}
	}

	result = database.DB.Raw(`select * from get_qtd_cards_por_status($1,$2,$3)`, IDProject, t1, t2).Find(&listCardsPerStatus)

	if result.Error != nil {
		return http.StatusBadRequest, gin.H{"error": "erro ao retornar as cards por tag"}
	}

	result = database.DB.Raw(`select * from get_retrabalhos($1,$2,$3)`, IDProject, t1, t2).Find(&listCardsRework)

	if result.Error != nil {
		return http.StatusBadRequest, gin.H{"error": "erro ao retornar a quantidade de retrabalho por card"}
	}

	result = database.DB.Raw(`select * from get_qtd_cards_finalizados_por_projeto($1,$2,$3)`, IDProject, t1, t2).Find(&listCardsFinished)

	if result.Error != nil {
		return http.StatusBadRequest, gin.H{"error": "erro ao retornar a quantidade cards finalizados"}
	}

	result = database.DB.Raw(`select * from get_tempo_execucao_por_card($1,$2,$3)`, IDProject, t1, t2).Find(&listCardsTimeExecution)

	if result.Error != nil {
		return http.StatusBadRequest, gin.H{"error": "erro ao retornar o tempo de execução dos cards"}
	}

	fmt.Println("tempo ", listCardsTimeExecution)

	response = gin.H{"success": statisticsdto.GetStatisticsResponse{TagData: listCardsPerTag, UserData: listCardsPerUser,
		StatusData: listCardsPerStatus, ReworkCards: listCardsRework, FinishedCards: listCardsFinished, ExecutionCards: listCardsTimeExecution}}

	return http.StatusOK, response
}
