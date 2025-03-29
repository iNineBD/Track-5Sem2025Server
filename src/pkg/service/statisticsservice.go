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

	result := database.DB.Raw(`select tag.name, sum(fato.qtd_card) as qtd_card_tag from dw_track.fato_card fato
	inner join dw_track.dim_tag tag on tag.id = fato.fk_id_tag where fato.fk_id_project = ?
	group by tag.name;`, idProject)

	if result.Error != nil {
		return http.StatusBadRequest, nil
	}

	if err := result.Scan(&listCardsPerTag).Error; err != nil {
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, listCardsPerTag
}

func GetCardsPerUser(idProject int) (status int, listCardsPerUser []statisticsdto.UserData) {

	err := utils.GetProject(int64(idProject))

	if err != nil {
		return http.StatusBadRequest, nil
	}

	result := database.DB.Raw(`select colaborador.full_name, sum(fato.qtd_card) as qtd_card_user from dw_track.fato_card fato
	inner join dw_track.dim_user colaborador on colaborador.id = fato.fk_id_user where fato.fk_id_project = ?
	group by colaborador.full_name;`, idProject)

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

	result := database.DB.Raw(`select status.name, sum(fato.qtd_card) as qtd_card_status from dw_track.fato_card fato
	inner join dw_track.dim_status status on status.id = fato.fk_id_status where fato.fk_id_project = ?
	group by status.name;`, idProject)

	if result.Error != nil {
		return http.StatusBadRequest, nil
	}

	if err := result.Scan(&listCardsPerStatus).Error; err != nil {
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, listCardsPerStatus
}
