package userdto

type AuthenticateRequest struct {
	Email       string `json:"email" binding:"required"`
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}
