package usermanagementdto

type RelationUserRoleResponse struct {
	IDUser   int64  `json:"id_user"`
	NameUser string `json:"name_user"`
	IDRole   int64  `json:"id_role"`
	NameRole string `json:"name_role"`
}
