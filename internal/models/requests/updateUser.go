package requests

type UpdateUserRequest struct {
	FirstName   *string `json:"first_name"`
	LastName    *string `json:"last_name"`
	MiddleName  *string `json:"middle_name"`
	Email       *string `json:"email"`
	OldPassword *string `json:"oldPassword"`
	NewPassword *string `json:"newPassword"`
}
