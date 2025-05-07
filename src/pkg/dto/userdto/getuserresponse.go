package userdto

type GetUserResponse struct {
	ID       int64  `json:"id_user,omitempty"`
	NameUser string `json:"name_user,omitempty"`
	Email    string `json:"email,omitempty"`
	Role     int64  `json:"id_role,omitempty"`
	NameRole string `json:"name_role,omitempty"`
	Token    string `json:"token,omitempty"`
}
