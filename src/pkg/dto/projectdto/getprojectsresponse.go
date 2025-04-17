package projectdto

type GetProjectsResponse struct {
	ID          int64  `json:"id_project"`
	Name        string `json:"name_project"`
	Description string `json:"description"`
}
