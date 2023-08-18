package domain

type Otp struct {
	IdOtp  int
	IdUser int
	Email  string
	Otp    string
	Time   int64
	Status string
}
