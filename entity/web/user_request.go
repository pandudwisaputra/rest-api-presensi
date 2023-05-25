package web

type CreateAkunRequest struct {
	Nama       string `validate:"required" json:"nama"`
	IdKaryawan string `validate:"required" json:"id_karyawan"`
	Email      string `validate:"required" json:"email"`
	NoHp       string `json:"no_hp"`
	Password   string `validate:"required" json:"password"`
}

type LoginRequest struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}

type ProfileRequest struct {
	IdUser int `validate:"required"`
}

type UpdatePasswordRequest struct {
	IdUser      int    `validate:"required" json:"id_user"`
	Email       string `validate:"required" json:"email"`
	NewPassword string `validate:"required" json:"new_password"`
}

type UpdateAvaRequest struct {
	IdUser int    `validate:"required" json:"id_user"`
	Ava    string `validate:"required" json:"ava"`
}
