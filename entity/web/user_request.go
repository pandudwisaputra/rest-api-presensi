package web

type CreateAkunRequest struct {
	IdKaryawan string `validate:"required" json:"id_karyawan"`
	Email      string `validate:"required" json:"email"`
	Password   string `validate:"required" json:"password"`
	AndroidId       string `validate:"required" json:"android_id"`
}

type CreateRecognitionRequest struct {
	IdUser         int    `validate:"required" json:"id_user"`
	Key            string `validate:"required" json:"key"`
	Name           string `validate:"required" json:"name"`
	LocationLeft   string `validate:"required" json:"location_left"`
	LocationTop    string `validate:"required" json:"location_top"`
	LocationRight  string `validate:"required" json:"location_right"`
	LocationBottom string `validate:"required" json:"location_bottom"`
	Embeddings     string `validate:"required" json:"embeddings"`
	Distance       string `validate:"required" json:"distance"`
}

type LoginRequest struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}
type SmartphoneCheckRequest struct {
	IdUser int    `validate:"required" json:"id_user"`
	AndroidId   string `validate:"required" json:"android_id"`
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
