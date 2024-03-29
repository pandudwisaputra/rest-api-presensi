package web

type OtpRequest struct {
	IdUser int    `json:"id_user"`
	Email  string `validate:"required" json:"email"`
}

type OtpValidationRequest struct {
	Email string `validate:"required" json:"email"`
	Otp   string `validate:"required" json:"otp"`
}
