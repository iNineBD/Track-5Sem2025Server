package statisticsdto

type GetStatisticsResponse struct {
	TagData        []TagData            `json:"card_tag"`
	UserData       []UserData           `json:"card_user"`
	StatusData     []StatusData         `json:"card_status"`
	ReworkCards    []ReworkCards        `json:"reworks_cards"`
	StartedCards   []StartedCards       `json:"started_cards"`
	FinishedCards  []FinishedCards      `json:"finished_cards"`
	ExecutionCards []TimeExecutionCards `json:"execution_cards"`
}
