package statisticsdto

type TimeExecutionCards struct {
	NameCard      string `json:"name_card" gorm:"column:name_card"`
	TimeExecution string `json:"time_execution" gorm:"column:tempo_execucao"`
}
