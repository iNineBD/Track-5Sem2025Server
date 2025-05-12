package service

import (
	"inine-track/pkg/database"
	"inine-track/pkg/dto/statisticsdto"
	"inine-track/pkg/models"
	"inine-track/pkg/service/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetMetrics(IDProject int64, data1 string, data2 string, idUser int64, idRole int64) (status int, response gin.H) {
	var role models.DimRole

	err := utils.GetProject(int64(IDProject))

	if err != nil {
		return http.StatusBadRequest, gin.H{"error": err.Error()}
	}

	result := database.DB.Where("id_role = ?", idRole).First(&role)

	if result.Error != nil {
		return http.StatusBadRequest, gin.H{"error": "erro ao trazer a role do usuário"}
	}

	var role_name string = role.NameRole

	t1, t2, err := utils.FormateDate(data1, data2)

	if err != nil {
		return http.StatusBadRequest, gin.H{"error": err}
	}

	role_name_upper := strings.ToUpper(role_name)

	if role_name_upper == "ADMIN" || role_name_upper == "GESTOR" {
		status, response := GetMetricsRole(IDProject, t1, t2, 0)

		return status, response
	} else {
		status, response := GetMetricsRole(IDProject, t1, t2, idUser)

		return status, response
	}
}

func GetMetricsRole(IDProject int64, data1 time.Time, data2 time.Time, idUser int64) (status int, response gin.H) {
	listCardsPerTag, err2 := GetListCardTags(IDProject, data1, data2, idUser)

	if err2 != nil {
		return http.StatusBadRequest, err2
	}

	listCardsPerUser, err2 := GetListCardsPerUser(IDProject, data1, data2, idUser)

	if err2 != nil {
		return http.StatusBadRequest, err2
	}

	listCardsPerStatus, err2 := GetListCardsPerStatus(IDProject, data1, data2, idUser)

	if err2 != nil {
		return http.StatusBadRequest, err2
	}

	listCardsRework, err2 := GetListCardsRework(IDProject, data1, data2, idUser)

	if err2 != nil {
		return http.StatusBadRequest, err2
	}

	listCardsStarted, err2 := GetListCardsStarted(IDProject, data1, data2, idUser)

	if err2 != nil {
		return http.StatusBadRequest, err2
	}

	listCardsFinished, err2 := GetListCardsFinished(IDProject, data1, data2, idUser)

	if err2 != nil {
		return http.StatusBadRequest, err2
	}

	listCardsTimeExecution, err2 := GetListCardsTimeExecution(IDProject, data1, data2, idUser)

	if err2 != nil {
		return http.StatusBadRequest, err2
	}

	response = gin.H{"success": statisticsdto.GetStatisticsResponse{TagData: listCardsPerTag, UserData: listCardsPerUser,
		StatusData: listCardsPerStatus, ReworkCards: listCardsRework, StartedCards: listCardsStarted, FinishedCards: listCardsFinished,
		ExecutionCards: listCardsTimeExecution}}

	return http.StatusOK, response
}

func GetListCardTags(IDProject int64, data1 time.Time, data2 time.Time, idUser int64) (listCardsPerTag []statisticsdto.TagData, err gin.H) {

	var result *gorm.DB

	if idUser == 0 {
		result = database.DB.Raw(`select * from get_qtd_cards_por_tag($1,$2,$3)`, IDProject, data1, data2).Find(&listCardsPerTag)
	} else {
		result = database.DB.Raw(`select * from get_qtd_cards_por_tag_operador($1,$2,$3,$4)`, IDProject, data1, data2, idUser).Find(&listCardsPerTag)
	}

	if result.Error != nil {
		return nil, gin.H{"error": "erro ao retornar a quantidade de cards por tag"}
	}

	return listCardsPerTag, nil
}

func GetListCardsPerUser(IDProject int64, data1 time.Time, data2 time.Time, idUser int64) (listCardsPerUser []statisticsdto.UserData, err gin.H) {

	var result *gorm.DB

	if idUser == 0 {
		result = database.DB.Raw(`select * from get_qtd_cards_por_colaborador($1,$2,$3)`, IDProject, data1, data2).Find(&listCardsPerUser)
	} else {
		result = database.DB.Raw(`select * from get_qtd_cards_por_colaborador_operador($1,$2,$3,$4)`, IDProject, data1, data2, idUser).Find(&listCardsPerUser)
	}

	if result.Error != nil {
		return nil, gin.H{"error": "erro ao retornar as cards por colaborador"}
	}

	return listCardsPerUser, nil
}

func GetListCardsPerStatus(IDProject int64, data1 time.Time, data2 time.Time, idUser int64) (listCardsPerStatus []statisticsdto.StatusData, err gin.H) {

	var result *gorm.DB

	if idUser == 0 {
		result = database.DB.Raw(`select * from get_qtd_cards_por_status($1,$2,$3)`, IDProject, data1, data2).Find(&listCardsPerStatus)
	} else {
		result = database.DB.Raw(`select * from get_qtd_cards_por_status_operador($1,$2,$3,$4)`, IDProject, data1, data2, idUser).Find(&listCardsPerStatus)
	}

	if result.Error != nil {
		return nil, gin.H{"error": "erro ao retornar as cards por status"}
	}

	return listCardsPerStatus, nil
}

func GetListCardsRework(IDProject int64, data1 time.Time, data2 time.Time, idUser int64) (listCardsRework []statisticsdto.ReworkCards, err gin.H) {

	var result *gorm.DB

	if idUser == 0 {
		result = database.DB.Raw(`select * from get_retrabalhos($1,$2,$3)`, IDProject, data1, data2).Find(&listCardsRework)
	} else {
		result = database.DB.Raw(`select * from get_retrabalhos_operador($1,$2,$3,$4)`, IDProject, data1, data2, idUser).Find(&listCardsRework)
	}

	if result.Error != nil {
		return nil, gin.H{"error": "erro ao retornar a quantidade de retrabalho por card"}
	}

	return listCardsRework, nil
}

func GetListCardsStarted(IDProject int64, data1 time.Time, data2 time.Time, idUser int64) (listCardsStarted []statisticsdto.StartedCards, err gin.H) {

	var result *gorm.DB

	if idUser == 0 {
		result = database.DB.Raw(`select * from get_qtd_cards_criados_por_projeto($1,$2,$3)`, IDProject, data1, data2).Find(&listCardsStarted)
	} else {
		result = database.DB.Raw(`select * from get_qtd_cards_criados_por_projeto_operador($1,$2,$3,$4)`, IDProject, data1, data2, idUser).Find(&listCardsStarted)
	}

	if result.Error != nil {
		return nil, gin.H{"error": "erro ao retornar a quantidade de cards criados por projeto"}
	}

	return listCardsStarted, nil
}

func GetListCardsFinished(IDProject int64, data1 time.Time, data2 time.Time, idUser int64) (listCardsFinished []statisticsdto.FinishedCards, err gin.H) {

	var result *gorm.DB

	if idUser == 0 {
		result = database.DB.Raw(`select * from get_qtd_cards_criados_por_projeto($1,$2,$3)`, IDProject, data1, data2).Find(&listCardsFinished)
	} else {
		result = database.DB.Raw(`select * from get_qtd_cards_criados_por_projeto_operador($1,$2,$3,$4)`, IDProject, data1, data2, idUser).Find(&listCardsFinished)
	}


	if result.Error != nil {
		return nil, gin.H{"error": "erro ao retornar a quantidade de cards finalizados por projet"}
	}

	return listCardsFinished, nil
}

func GetListCardsTimeExecution(IDProject int64, data1 time.Time, data2 time.Time, idUser int64) (listCardsTimeExecution []statisticsdto.TimeExecutionCards, err gin.H) {

	var result *gorm.DB

	if idUser == 0 {
		result = database.DB.Raw(`select * from get_tempo_execucao_por_card($1,$2,$3)`, IDProject, data1, data2).Find(&listCardsTimeExecution)
	} else {
		result = database.DB.Raw(`select * from get_tempo_execucao_por_card_operador($1,$2,$3,$4)`, IDProject, data1, data2, idUser).Find(&listCardsTimeExecution)
	}

	if result.Error != nil {
		return nil, gin.H{"error": "erro ao retornar o tempo de execução dos cards"}
	}

	return listCardsTimeExecution, nil
}
