package service

import (
	"inine-track/pkg/database"
	"inine-track/pkg/dto/statisticsdto"
	"inine-track/pkg/service/utils"
	"net/http"
)

func GetCardsPerTag(idProject int) (status int, listCardsPerTag []statisticsdto.TagData) {

	err := utils.GetProject(int64(idProject))

	if err != nil {
		return http.StatusBadRequest, nil
	}

	result := database.DB.Raw(`select tag.name_tag, sum(fato.qtd_cards) as qtd_card_tag from dw_track.fato_cards fato
	inner join dw_track.dim_tag tag on tag.id_tag = fato.id_tag where fato.id_project = ?
	group by tag.name_tag;`, idProject)

	if result.Error != nil {
		return http.StatusBadRequest, nil
	}

	if err := result.Scan(&listCardsPerTag).Error; err != nil {
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, listCardsPerTag
}

func GetCardsPerUser(IDProject int) (status int, listCardsPerUser []statisticsdto.UserData) {

	err := utils.GetProject(int64(IDProject))

	if err != nil {
		return http.StatusBadRequest, nil
	}

	result := database.DB.Raw(`select colaborador.name_user, sum(fato.qtd_cards) as qtd_card_user from dw_track.fato_cards fato
	inner join dw_track.dim_user colaborador on colaborador.id_user = fato.id_user where fato.id_project = ?
	group by colaborador.name_user;`, IDProject)

	if result.Error != nil {
		return http.StatusBadRequest, nil
	}

	if err := result.Scan(&listCardsPerUser).Error; err != nil {
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, listCardsPerUser
}

func GetCardsPerStatus(idProject int) (status int, listCardsPerStatus []statisticsdto.StatusData) {

	err := utils.GetProject(int64(idProject))

	if err != nil {
		return http.StatusBadRequest, nil
	}

	result := database.DB.Raw(`select status.name_status, sum(fato.qtd_cards) as qtd_card_status from dw_track.fato_cards fato
	inner join dw_track.dim_status status on status.id_status = fato.id_status where fato.id_project = ?
	group by status.name_status;`, idProject)

	if result.Error != nil {
		return http.StatusBadRequest, nil
	}

	if err := result.Scan(&listCardsPerStatus).Error; err != nil {
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, listCardsPerStatus
}
