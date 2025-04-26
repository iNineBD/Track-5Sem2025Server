package statisticsdto

type StartedCards struct {
	NameProject string `json:"name_project" gorm:"column:name_project"`
	Qtd         int64  `json:"qtd_cards_started" gorm:"column:qtd_card_started"`
}
