package service

import (
	"inine-track/pkg/database"
	"inine-track/pkg/dto/statisticsdto"
	"inine-track/pkg/service/utils"
	"net/http"
)

func GetCardsPerTag(IDProject int, data1 string, data2 string) (status int, listCardsPerTag []statisticsdto.TagData) {

	err := utils.GetProject(int64(IDProject))

	if err != nil {
		return http.StatusBadRequest, nil
	}

	t1, t2, err := utils.FormateDate(data1, data2)

	if err != nil {
		return http.StatusBadRequest, nil
	}

	result := database.DB.Raw(`select * from get_qtd_cards_por_tag($1,$2,$3)`, IDProject, t1, t2)

	if result.Error != nil {
		return http.StatusBadRequest, nil
	}

	if err := result.Scan(&listCardsPerTag).Error; err != nil {
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, listCardsPerTag
}

func GetCardsPerUser(IDProject int, data1 string, data2 string) (status int, listCardsPerUser []statisticsdto.UserData) {

	err := utils.GetProject(int64(IDProject))

	if err != nil {
		return http.StatusBadRequest, nil
	}

	t1, t2, err := utils.FormateDate(data1, data2)

	if err != nil {
		return http.StatusBadRequest, nil
	}

	result := database.DB.Raw(`select * from get_qtd_cards_por_colaborador($1,$2,$3)`, IDProject, t1, t2)

	if result.Error != nil {
		return http.StatusBadRequest, nil
	}

	if err := result.Scan(&listCardsPerUser).Error; err != nil {
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, listCardsPerUser
}

func GetCardsPerStatus(IDProject int, data1 string, data2 string) (status int, listCardsPerStatus []statisticsdto.StatusData) {

	err := utils.GetProject(int64(IDProject))

	if err != nil {
		return http.StatusBadRequest, nil
	}

	t1, t2, err := utils.FormateDate(data1, data2)

	if err != nil {
		return http.StatusBadRequest, nil
	}

	result := database.DB.Raw(`select * from get_qtd_cards_por_status($1,$2,$3)`, IDProject, t1, t2)

	if result.Error != nil {
		return http.StatusBadRequest, nil
	}

	if err := result.Scan(&listCardsPerStatus).Error; err != nil {
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, listCardsPerStatus
}

func GetTimeDesenvCards(IDProject int, data1 string, data2 string) (status int, listCardsRework []statisticsdto.ReworkCards) {

	err := utils.GetProject(int64(IDProject))

	if err != nil {
		return http.StatusBadRequest, nil
	}

	t1, t2, err := utils.FormateDate(data1, data2)

	if err != nil {
		return http.StatusBadRequest, nil
	}

	result := database.DB.Raw(`select * from get_retrabalhos($1,$2,$3)`, IDProject, t1, t2)

	if result.Error != nil {
		return http.StatusBadRequest, nil
	}

	if err := result.Scan(&listCardsRework).Error; err != nil {
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, listCardsRework
}
