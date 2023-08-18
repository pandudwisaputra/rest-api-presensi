package web

type PresensiMasukResponse struct {
	IdPresensi int    `json:"id_presensi"`
	Keterangan string `json:"keterangan"`
}

type PresensiTidakMasukResponse struct {
	IdPresensi int    `json:"id_presensi"`
	Keterangan string `json:"keterangan"`
}

type PresensiKeluarResponse struct {
	Keterangan string `json:"keterangan"`
}

type RiwayatPresensiResponse struct {
	IdPresensi           int    `json:"id_presensi"`
	IdUser               int    `json:"id_user"`
	TanggalPresensi      string `json:"tanggal_presensi"`
	JamMasuk             string `json:"jam_masuk"`
	JamPulang            string `json:"jam_pulang"`
	TanggalPulang        string `json:"tanggal_pulang"`
	KeteranganMasuk      string `json:"keterangan_masuk"`
	KeteranganKeluar     string `json:"keterangan_keluar"`
	Latitude             string `json:"latitude"`
	Longitude            string `json:"longitude"`
	Alamat               string `json:"alamat"`
	Status               string `json:"status"`
	KeteranganTidakMasuk string `json:"keterangan_tidak_masuk"`
	LinkBukti            string `json:"link_bukti"`
}

type PresensiCheckResponse struct {
	IdPresensi       int    `json:"id_presensi"`
	IdUser           int    `json:"id_user"`
	TanggalPresensi  string `json:"tanggal_presensi"`
	JamMasuk         string `json:"jam_masuk"`
	JamPulang        string `json:"jam_pulang"`
	TanggalPulang    string `json:"tanggal_pulang"`
	KeteranganMasuk  string `json:"keterangan_masuk"`
	KeteranganKeluar string `json:"keterangan_keluar"`
	Latitude         string `json:"latitude"`
	Longitude        string `json:"longitude"`
	Alamat           string `json:"alamat"`
	Status           string `json:"status"`
}
