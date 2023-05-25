package web

type PresensiMasukRequest struct {
	IdUser          int    `validate:"required" json:"id_user"`
	TanggalPresensi string `validate:"required" json:"tanggal_presensi"`
	Latitude        string `validate:"required" json:"latitude"`
	Longitude       string `validate:"required" json:"longitude"`
	Selfie          string `validate:"required" json:"selfie"`
	Alamat          string `validate:"required" json:"alamat"`
}

type PresensiKeluarRequest struct {
	IdUser           int    `validate:"required" json:"id_user"`
	TanggalPresesnsi string `validate:"required" json:"tanggal_presesnsi"`
}

type RiwayatPresensiRequest struct {
	IdUser int `validate:"required" json:"id_user"`
}
