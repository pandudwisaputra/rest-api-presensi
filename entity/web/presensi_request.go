package web

type PresensiMasukRequest struct {
	IdUser          int    `validate:"required" json:"id_user"`
	TanggalPresensi string `validate:"required" json:"tanggal_presensi"`
	Latitude        string `validate:"required" json:"latitude"`
	Longitude       string `validate:"required" json:"longitude"`
	Selfie          string `validate:"required" json:"selfie"`
	Alamat          string `validate:"required" json:"alamat"`
}

type PresensiTidakMasukRequest struct {
	IdUser               int    `validate:"required" json:"id_user"`
	TanggalPresensi      string `validate:"required" json:"tanggal_presensi"`
	KeteranganTidakMasuk string `validate:"required" json:"keterangan_tidak_masuk"`
	LinkBukti            string `validate:"required" json:"link_bukti"`
}

type PresensiKeluarRequest struct {
	IdUser          int    `validate:"required" json:"id_user"`
	IdPresensi      int    `validate:"required" json:"id_presensi"`
	TanggalPresensi string `validate:"required" json:"tanggal_presensi"`
}

type RiwayatPresensiRequest struct {
	IdUser int `validate:"required" json:"id_user"`
}
