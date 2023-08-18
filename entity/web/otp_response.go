package web

type OtpResponse struct {
	Email  string `json:"email"`
	Status string `json:"status"`
}

type OtpValidationResponse struct {
	Email  string `json:"email"`
	Otp    string `json:"otp"`
	Status string `json:"status"`
}
