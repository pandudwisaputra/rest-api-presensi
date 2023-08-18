package web

type CreateAkunResponse struct {
	IdUser int    `json:"id_user"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

type CreateRecognitionResponse struct {
	IdRecognition int    `json:"id_recognition"`
	Status        string `json:"status"`
}

type LoginResponse struct {
	IdUser int    `json:"id_user"`
	Email  string `json:"email"`
	Status string `json:"status"`
}
type GetSmartphoneCheckResponse struct {
	IdUser int    `json:"id_user"`
	AndroidId  string `json:"android_id"`
}

type GetIdUserResponse struct {
	IdUser int    `json:"id_user"`
	Email  string `json:"email"`
}

type UpdatePasswordResponse struct {
	Email  string `json:"email"`
	Status string `json:"status"`
}

type UpdateAvaResponse struct {
	IdUser int    `json:"id_user"`
	Status string `json:"status"`
}

type GetProfileResponse struct {
	IdUser     int    `json:"id_user"`
	IdKaryawan string `json:"id_karyawan"`
	Email      string `json:"email"`
	Avatar     string `json:"avatar"`
	Nama       string `json:"nama"`
	NoHp       string `json:"no_hp"`
}

type GetRecognitionResponse struct {
	IdRecognition  int    `json:"id_recognition"`
	IdUser         int    `json:"id_user"`
	Key            string `json:"key"`
	Name           string `json:"name"`
	LocationLeft   string `json:"location_left"`
	LocationTop    string `json:"location_top"`
	LocationRight  string `json:"location_right"`
	LocationBottom string `json:"location_bottom"`
	Embeddings     string `json:"embeddings"`
	Distance       string `json:"distance"`
}

type GetEmailCheckResponse struct {
	Status string `json:"status"`
}

type GetKaryawanResponse struct {
	IdKaryawan  string `json:"id_karyawan"`
	IdJabatan   int    `json:"id_jabatan"`
	NamaLengkap string `json:"nama_lengkap"`
	Alamat      string `json:"alamat"`
	Agama       string `json:"agama"`
	Email       string `json:"email"`
	NoHp        string `json:"no_hp"`
	Pendidikan  string `json:"pendidikan"`
	Jabatan     string `json:"jabatan"`
}
type GetStatusKaryawanResponse struct {
	StatusKaryawan  string `json:"status_karyawan"`
}