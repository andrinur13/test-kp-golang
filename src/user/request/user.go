package request

type RegisterUserRequest struct {
	FullName          string `json:"full_name"`
	LegalName         string `json:"legal_name"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	BornCity          string `json:"born_city"`
	BornDate          string `json:"born_date"`
	Income            int    `json:"income"`
	IdentityPhotoPath string `json:"identity_photo_path"`
	SelfiePhotoPath   string `json:"selfie_photo_path"`
}
