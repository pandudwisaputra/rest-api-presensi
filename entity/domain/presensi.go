package domain

type Presensi struct {
	IdPresensi       int
	IdUser           int
	TanggalPresensi  string
	JamMasuk         string
	JamPulang        string
	TanggalPulang    string
	KeteranganMasuk  string
	KeteranganKeluar string
	Latitude         string
	Longitude        string
	Alamat           string
	Status           string
	KeteranganTidakMasuk             string
	LinkBukti        string
}
