package service

import (
	"inine-track/pkg/database"
	"inine-track/pkg/dto/statisticsdto"
	"inine-track/pkg/service/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetMetrics(IDProject int64, data1 string, data2 string) (status int, response gin.H) {

	err := utils.GetProject(int64(IDProject))

	if err != nil {
		return http.StatusBadRequest, gin.H{"error": err.Error()}
	}

	t1, t2, err := utils.FormateDate(data1, data2)

	if err != nil {
		return http.StatusBadRequest, gin.H{"error": err}
	}

	listCardsPerTag, err2 := GetListCardTags(IDProject, t1, t2)

	if err2 != nil {
		return http.StatusBadRequest, err2
	}

	listCardsPerUser, err2 := GetListCardsPerUser(IDProject, t1, t2)

	if err2 != nil {
		return http.StatusBadRequest, err2
	}

	listCardsPerStatus, err2 := GetListCardsPerStatus(IDProject, t1, t2)

	if err2 != nil {
		return http.StatusBadRequest, err2
	}

	listCardsRework, err2 := GetListCardsRework(IDProject, t1, t2)

	if err2 != nil {
		return http.StatusBadRequest, err2
	}

	listCardsStarted, err2 := GetListCardsStarted(IDProject, t1, t2)

	if err2 != nil {
		return http.StatusBadRequest, err2
	}

	listCardsFinished, err2 := GetListCardsFinished(IDProject, t1, t2)

	if err2 != nil {
		return http.StatusBadRequest, err2
	}

	listCardsTimeExecution, err2 := GetListCardsTimeExecution(IDProject, t1, t2)

	if err2 != nil {
		return http.StatusBadRequest, err2
	}

	response = gin.H{"success": statisticsdto.GetStatisticsResponse{TagData: listCardsPerTag, UserData: listCardsPerUser,
		StatusData: listCardsPerStatus, ReworkCards: listCardsRework, StartedCards: listCardsStarted, FinishedCards: listCardsFinished,
		ExecutionCards: listCardsTimeExecution}}

	return http.StatusOK, response
}

func GetListCardTags(IDProject int64, data1 time.Time, data2 time.Time) (listCardsPerTag []statisticsdto.TagData, err gin.H) {

	result := database.DB.Raw(`select * from get_qtd_cards_por_tag($1,$2,$3)`, IDProject, data1, data2).Find(&listCardsPerTag)

	if result.Error != nil {
		return nil, gin.H{"error": "erro ao retornar a quantidade de cards por tag"}
	}

	return listCardsPerTag, nil
}

func GetListCardsPerUser(IDProject int64, data1 time.Time, data2 time.Time) (listCardsPerUser []statisticsdto.UserData, err gin.H) {

	result := database.DB.Raw(`select * from get_qtd_cards_por_colaborador($1,$2,$3)`, IDProject, data1, data2).Find(&listCardsPerUser)

	if result.Error != nil {
		return nil, gin.H{"error": "erro ao retornar as cards por colaborador"}
	}

	return listCardsPerUser, nil
}

func GetListCardsPerStatus(IDProject int64, data1 time.Time, data2 time.Time) (listCardsPerStatus []statisticsdto.StatusData, err gin.H) {

	result := database.DB.Raw(`select * from get_qtd_cards_por_status($1,$2,$3)`, IDProject, data1, data2).Find(&listCardsPerStatus)

	if result.Error != nil {
		return nil, gin.H{"error": "erro ao retornar as cards por status"}
	}

	return listCardsPerStatus, nil
}

func GetListCardsRework(IDProject int64, data1 time.Time, data2 time.Time) (listCardsRework []statisticsdto.ReworkCards, err gin.H) {

	result := database.DB.Raw(`select * from get_retrabalhos($1,$2,$3)`, IDProject, data1, data2).Find(&listCardsRework)

	if result.Error != nil {
		return nil, gin.H{"error": "erro ao retornar a quantidade de retrabalho por card"}
	}

	return listCardsRework, nil
}

func GetListCardsStarted(IDProject int64, data1 time.Time, data2 time.Time) (listCardsStarted []statisticsdto.StartedCards, err gin.H) {

	result := database.DB.Raw(`select * from get_qtd_cards_criados_por_projeto($1,$2,$3)`, IDProject, data1, data2).Find(&listCardsStarted)

	if result.Error != nil {
		return nil, gin.H{"error": "erro ao retornar a quantidade de cards criados por projeto"}
	}

	return listCardsStarted, nil
}

func GetListCardsFinished(IDProject int64, data1 time.Time, data2 time.Time) (listCardsFinished []statisticsdto.FinishedCards, err gin.H) {

	result := database.DB.Raw(`select * from get_qtd_cards_criados_por_projeto($1,$2,$3)`, IDProject, data1, data2).Find(&listCardsFinished)

	if result.Error != nil {
		return nil, gin.H{"error": "erro ao retornar a quantidade de cards finalizados por projet"}
	}

	return listCardsFinished, nil
}

func GetListCardsTimeExecution(IDProject int64, data1 time.Time, data2 time.Time) (listCardsTimeExecution []statisticsdto.TimeExecutionCards, err gin.H) {

	result := database.DB.Raw(`select * from get_qtd_cards_criados_por_projeto($1,$2,$3)`, IDProject, data1, data2).Find(&listCardsTimeExecution)

	if result.Error != nil {
		return nil, gin.H{"error": "erro ao retornar o tempo de execução dos cards"}
	}

	return listCardsTimeExecution, nil
}
