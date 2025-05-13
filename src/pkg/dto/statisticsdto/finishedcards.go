package statisticsdto

type FinishedCards struct {
	NameProject string `json:"name_project" gorm:"column:name_project"`
	Qtd         int64  `json:"qtd_cards_finished" gorm:"column:qtd_card_finished"`
}
