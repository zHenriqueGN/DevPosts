package models

// PasswordUpdate represents an object used to update passwords
type PasswordUpdate struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
