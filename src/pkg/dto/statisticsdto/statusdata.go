package statisticsdto

type StatusData struct {
	Status string `json:"status" gorm:"column:name"`
	Qtd    int64  `json:"qtd" gorm:"column:qtd_card_status"`
}
