package helper

import (
	"golang/rest-api-presensi/entity/domain"
	"golang/rest-api-presensi/entity/web"
)

func ToCreateAkunResponse(user domain.User) web.CreateAkunResponse {
	return web.CreateAkunResponse{
		IdUser: user.IdUser,
		Email:  user.Email,
		Status: "Success",
	}
}

func ToCreateRecognitionResponse(recognition domain.Recognition) web.CreateRecognitionResponse {
	return web.CreateRecognitionResponse{
		IdRecognition: recognition.IdRecognition,
		Status:        "Success",
	}
}

func ToOtpResponse(otp domain.Otp) web.OtpResponse {
	return web.OtpResponse{
		Email:  otp.Email,
		Status: "OTP Terkirim",
	}
}

func ToOtpValidationResponse(otp domain.Otp) web.OtpValidationResponse {
	return web.OtpValidationResponse{
		Email:  otp.Email,
		Status: otp.Status,
	}
}

func ToUpdatePasswordResponse(user domain.User) web.UpdatePasswordResponse {
	return web.UpdatePasswordResponse{
		Email:  user.Email,
		Status: "success",
	}
}

func ToLoginResponse(nasabah domain.User) web.LoginResponse {
	return web.LoginResponse{
		IdUser: nasabah.IdUser,
		Email:  nasabah.Email,
		Status: nasabah.Status,
	}
}

func ToGetIdUserResponse(nasabah domain.User) web.GetIdUserResponse {
	return web.GetIdUserResponse{
		IdUser: nasabah.IdUser,
		Email:  nasabah.Email,
	}
}

func ToUpdateAvaResponse(user domain.User) web.UpdateAvaResponse {
	return web.UpdateAvaResponse{
		IdUser: user.IdUser,
		Status: "update avatar berhasil",
	}
}

func ToGetProfileResponse(user domain.User) web.GetProfileResponse {
	return web.GetProfileResponse{
		IdUser:     user.IdUser,
		IdKaryawan: user.IdKaryawan,
		Email:      user.Email,
		Avatar:     user.Avatar,
		Nama:       user.NamaLengkap,
		NoHp:       user.NoHp,
	}
}

func ToGetRecognitionResponse(recognition domain.Recognition) web.GetRecognitionResponse {
	return web.GetRecognitionResponse{
		IdRecognition:  recognition.IdRecognition,
		IdUser:         recognition.IdUser,
		Key:            recognition.Key,
		Name:           recognition.Name,
		LocationLeft:   recognition.LocationLeft,
		LocationTop:    recognition.LocationTop,
		LocationRight:  recognition.LocationRight,
		LocationBottom: recognition.LocationBottom,
		Embeddings:     recognition.Embeddings,
		Distance:       recognition.Distance,
	}
}

func ToGetEmailResponse(string2 string) web.GetEmailCheckResponse {
	return web.GetEmailCheckResponse{
		Status: string2,
	}
}

func ToGetSmartphoneCheckResponse(smartphoneCheck domain.User) web.GetSmartphoneCheckResponse {
	return web.GetSmartphoneCheckResponse{
		IdUser: smartphoneCheck.IdUser,
		AndroidId: smartphoneCheck.AndroidId,
	}
}

func ToGetKaryawanResponse(karyawan domain.Karyawan) web.GetKaryawanResponse {
	return web.GetKaryawanResponse{
		IdKaryawan:  karyawan.IdKaryawan,
		IdJabatan:   karyawan.IdJabatan,
		NamaLengkap: karyawan.NamaLengkap,
		Alamat:      karyawan.Alamat,
		Agama:       karyawan.Agama,
		Email:       karyawan.Email,
		NoHp:        karyawan.NoHp,
		Pendidikan:  karyawan.Pendidikan,
		Jabatan:     karyawan.Jabatan,
	}
}
func ToGetStatusKaryawanResponse(karyawan domain.Karyawan) web.GetStatusKaryawanResponse {
	return web.GetStatusKaryawanResponse{
		StatusKaryawan: karyawan.StatusKaryawan,
	}
}

func ToGetOfficeDataResponse(office domain.Office) web.OfficeResponse {
	return web.OfficeResponse{
		NamaKantor: office.NamaKantor,
		Alamat:     office.Alamat,
		Latitude:   office.Latitude,
		Longitude:  office.Longitude,
		JamMasuk:   office.JamMasuk,
		JamPulang:  office.JamPulang,
		Radius:     office.Radius,
	}
}

func ToPresensiMasukResponse(presensi domain.Presensi) web.PresensiMasukResponse {
	return web.PresensiMasukResponse{
		IdPresensi: presensi.IdPresensi,
		Keterangan: presensi.KeteranganMasuk,
	}

}

func ToPresensiTidakMasukResponse(presensi domain.Presensi) web.PresensiTidakMasukResponse {
	return web.PresensiTidakMasukResponse{
		IdPresensi: presensi.IdPresensi,
		Keterangan: presensi.KeteranganMasuk,
	}

}

func ToPresensiKeluarResponse(presensi domain.Presensi) web.PresensiKeluarResponse {
	return web.PresensiKeluarResponse{
		Keterangan: presensi.KeteranganKeluar,
	}

}

func ToRiwayatPresensiResponse(presensi domain.Presensi) web.RiwayatPresensiResponse {
	return web.RiwayatPresensiResponse{
		IdPresensi:           presensi.IdPresensi,
		IdUser:               presensi.IdUser,
		TanggalPresensi:      presensi.TanggalPresensi,
		JamMasuk:             presensi.JamMasuk,
		JamPulang:            presensi.JamPulang,
		TanggalPulang:        presensi.TanggalPulang,
		KeteranganMasuk:      presensi.KeteranganMasuk,
		KeteranganKeluar:     presensi.KeteranganKeluar,
		Latitude:             presensi.Latitude,
		Longitude:            presensi.Longitude,
		Alamat:               presensi.Alamat,
		Status:               presensi.Status,
		KeteranganTidakMasuk: presensi.KeteranganTidakMasuk,
		LinkBukti:            presensi.LinkBukti,
	}
}

func ToRiwayatPresensiResponses(presensi []domain.Presensi) []web.RiwayatPresensiResponse {
	var riwayatResponses []web.RiwayatPresensiResponse
	for _, presensi := range presensi {
		riwayatResponses = append(riwayatResponses, ToRiwayatPresensiResponse(presensi))
	}
	return riwayatResponses
}

func ToPresensiCheckResponse(presensi domain.Presensi) web.PresensiCheckResponse {
	return web.PresensiCheckResponse{
		IdPresensi:       presensi.IdPresensi,
		IdUser:           presensi.IdUser,
		TanggalPresensi:  presensi.TanggalPresensi,
		JamMasuk:         presensi.JamMasuk,
		JamPulang:        presensi.JamPulang,
		TanggalPulang:    presensi.TanggalPulang,
		KeteranganMasuk:  presensi.KeteranganMasuk,
		KeteranganKeluar: presensi.KeteranganKeluar,
		Latitude:         presensi.Latitude,
		Longitude:        presensi.Longitude,
		Alamat:           presensi.Alamat,
		Status:           presensi.Status,
	}
}
