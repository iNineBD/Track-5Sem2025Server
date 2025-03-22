package statisticsdto

type GetStatisticsResponse struct {
	TagData    []TagData    `json:"card_tag"`
	UserData   []UserData   `json:"card_user"`
	StatusData []StatusData `json:"card_status"`
}
