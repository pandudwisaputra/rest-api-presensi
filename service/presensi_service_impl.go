package service

import (
	"database/sql"
	"golang/rest-api-presensi/entity/domain"
	"golang/rest-api-presensi/entity/web"
	"golang/rest-api-presensi/exception"
	"golang/rest-api-presensi/helper"
	"golang/rest-api-presensi/repository"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/net/context"
)

type PresensiServiceImpl struct {
	PresensiRepository repository.PresensiRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewPresensiServiceImpl(presensiRepository repository.PresensiRepository, DB *sql.DB, validate *validator.Validate) *PresensiServiceImpl {
	return &PresensiServiceImpl{
		PresensiRepository: presensiRepository,
		DB:                 DB,
		Validate:           validate}
}

func (service *PresensiServiceImpl) PresensiMasuk(ctx context.Context, request web.PresensiMasukRequest) web.PresensiMasukResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	office := helper.Presensi{}
	office = helper.GetOffice()

	parts := strings.Split(office.JamMasuk, ":")
	hour, err := strconv.Atoi(parts[0])
	minute, err := strconv.Atoi(parts[1])

	//now := time.Now()

	loc, err := time.LoadLocation("Asia/Jakarta") // mengatur zona waktu ke WIB
	helper.PanicIfError(err)
	now := time.Now().In(loc)
	nowhour := now.Hour()
	nowminute := now.Minute()

	var status string
	if nowhour >= hour && nowminute > minute {
		status = "Telat"
	} else {
		status = "Tepat Waktu"
	}

	null := "-"
	presensi := domain.Presensi{
		IdUser:               request.IdUser,
		TanggalPresensi:      request.TanggalPresensi,
		JamMasuk:             strconv.Itoa(nowhour) + ":" + strconv.Itoa(nowminute),
		KeteranganMasuk:      status,
		Latitude:             request.Latitude,
		Longitude:            request.Longitude,
		Alamat:               request.Alamat,
		JamPulang:            null,
		TanggalPulang:        null,
		Status:               null,
		KeteranganKeluar:     null,
		KeteranganTidakMasuk: null,
		LinkBukti:            null,
	}

	presensi = service.PresensiRepository.PresensiMasuk(ctx, tx, presensi)

	return helper.ToPresensiMasukResponse(presensi)
}

func (service *PresensiServiceImpl) PresensiTidakMasuk(ctx context.Context, request web.PresensiTidakMasukRequest) web.PresensiTidakMasukResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	//now := time.Now()
	loc, err := time.LoadLocation("Asia/Jakarta") // mengatur zona waktu ke WIB

	helper.PanicIfError(err)
	now := time.Now().In(loc)
	nowhour := now.Hour()
	nowminute := now.Minute()

	null := "-"
	presensi := domain.Presensi{
		IdUser:               request.IdUser,
		TanggalPresensi:      request.TanggalPresensi,
		JamMasuk:             strconv.Itoa(nowhour) + "." + strconv.Itoa(nowminute),
		KeteranganMasuk:      "Tidak Masuk",
		Latitude:             null,
		Longitude:            null,
		Alamat:               null,
		JamPulang:            null,
		TanggalPulang:        null,
		Status:               "Selesai",
		KeteranganKeluar:     null,
		KeteranganTidakMasuk: request.KeteranganTidakMasuk,
		LinkBukti:            request.LinkBukti,
	}

	presensi = service.PresensiRepository.PresensiTidakMasuk(ctx, tx, presensi)

	return helper.ToPresensiTidakMasukResponse(presensi)
}

func (service *PresensiServiceImpl) PresensiKeluar(ctx context.Context, request web.PresensiKeluarRequest) web.PresensiKeluarResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	office := helper.Presensi{}
	office = helper.GetOffice()

	parts := strings.Split(office.JamPulang, ":")
	hour, err := strconv.Atoi(parts[0])
	helper.PanicIfError(err)
	minute, err := strconv.Atoi(parts[1])
	helper.PanicIfError(err)

	loc, err := time.LoadLocation("Asia/Jakarta") // mengatur zona waktu ke WIB
	helper.PanicIfError(err)
	now := time.Now().In(loc)
	nowhour := now.Hour()
	nowminute := now.Minute()
	var status string
	if nowhour >= hour && nowminute > minute {
		status = "Pulang Tepat Waktu"
	} else {
		status = "Pulang Lebih Awal"
	}

	waktu := time.Now().UnixNano() / 1000000
	presensi := domain.Presensi{
		IdUser:           request.IdUser,
		IdPresensi:       request.IdPresensi,
		TanggalPresensi:  request.TanggalPresensi,
		TanggalPulang:    strconv.Itoa(int(waktu)),
		JamPulang:        strconv.Itoa(nowhour) + ":" + strconv.Itoa(nowminute),
		KeteranganKeluar: status,
		Status:           "Selesai",
	}

	presensi = service.PresensiRepository.PresensiKeluar(ctx, tx, presensi)

	return helper.ToPresensiKeluarResponse(presensi)
}

func (service *PresensiServiceImpl) Riwayat(ctx context.Context, request int) []web.RiwayatPresensiResponse {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	riwayatPresensi, err := service.PresensiRepository.Riwayat(ctx, tx, request)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToRiwayatPresensiResponses(riwayatPresensi)
}

func (service *PresensiServiceImpl) PresensiCheck(ctx context.Context, request int) web.PresensiCheckResponse {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	presensiCheck, err := service.PresensiRepository.PresensiCheck(ctx, tx, request)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToPresensiCheckResponse(presensiCheck)
}
