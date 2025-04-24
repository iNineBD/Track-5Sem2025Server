package userdto


type GetUserResponse struct {
	ID int64 `json:"id_user,omitempty"`
	Email string `json:"email,omitempty"`
	Role int64 `json:"id_role,omitempty"`.
	Token string `json:"token,omitempty"`
}
