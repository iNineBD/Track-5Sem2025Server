package statisticsdto

type ReworkCards struct {
	CardName  string `json:"card_name" gorm:"column:name_card"`
	QtdRework int64  `json:"qtd_rework" gorm:"column:qtd_retrabalhos"`
}
