package statisticsdto

type GetStatisticsResponse struct {
	TagData    []TagData    `json:"tag_data"`
	UserData   []UserData   `json:"user_data"`
	StatusData []StatusData `json:"status_data"`
}
