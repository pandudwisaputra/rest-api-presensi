package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang/rest-api-presensi/entity/domain"
	"golang/rest-api-presensi/helper"
)

type PresensiRepositoryImpl struct {
}

func NewPresensiRepositoryImpl() *PresensiRepositoryImpl {
	return &PresensiRepositoryImpl{}
}

func (repository *PresensiRepositoryImpl) PresensiMasuk(ctx context.Context, tx *sql.Tx, presensi domain.Presensi) domain.Presensi {
	script := "insert into presensi_masuk(id_user, tanggal_presensi, jam_masuk, jam_keluar, tanggal_keluar, keterangan_masuk, keterangan_keluar, latitude, longitude, selfie, alamat, status_presensi, keterangan_tidak_masuk, link_bukti) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, script, presensi.IdUser, presensi.TanggalPresensi, presensi.JamMasuk, presensi.JamPulang, presensi.TanggalPulang, presensi.KeteranganMasuk, presensi.KeteranganKeluar, presensi.Latitude, presensi.Longitude, presensi.Selfie, presensi.Alamat, presensi.Status, presensi.KeteranganTidakMasuk, presensi.LinkBukti)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	presensi.IdUser = int(id)

	return presensi
}

func (repository *PresensiRepositoryImpl) PresensiTidakMasuk(ctx context.Context, tx *sql.Tx, presensi domain.Presensi) domain.Presensi {
	script := "insert into presensi_masuk(id_user, tanggal_presensi, jam_masuk, jam_keluar, tanggal_keluar, keterangan_masuk, keterangan_keluar, latitude, longitude, selfie, alamat, status_presensi, keterangan_tidak_masuk, link_bukti) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, script, presensi.IdUser, presensi.TanggalPresensi, presensi.JamMasuk, presensi.JamPulang, presensi.TanggalPulang, presensi.KeteranganMasuk, presensi.KeteranganKeluar, presensi.Latitude, presensi.Longitude, presensi.Selfie, presensi.Alamat, presensi.Status, presensi.KeteranganTidakMasuk, presensi.LinkBukti)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	presensi.IdUser = int(id)

	return presensi
}

func (repository *PresensiRepositoryImpl) PresensiKeluar(ctx context.Context, tx *sql.Tx, presensi domain.Presensi) domain.Presensi {
	script := "update presensi_masuk set jam_keluar = ?, tanggal_keluar = ?,keterangan_keluar = ? ,  status_presensi = ? where id_user = ? and tanggal_presensi = ?"
	_, err := tx.ExecContext(ctx, script, presensi.JamPulang, presensi.TanggalPulang, presensi.KeteranganKeluar, presensi.Status, presensi.IdUser, presensi.TanggalPresensi)
	helper.PanicIfError(err)
	return presensi
}

func (repository *PresensiRepositoryImpl) Riwayat(ctx context.Context, tx *sql.Tx, presensi int) ([]domain.Presensi, error) {
	script := "select id_presensi, id_user, tanggal_presensi, jam_masuk, jam_keluar, tanggal_keluar, keterangan_masuk, keterangan_keluar, latitude, longitude, alamat, status_presensi from presensi_masuk where id_user = ?"
	rows, err := tx.QueryContext(ctx, script, presensi)
	helper.PanicIfError(err)
	defer rows.Close()

	var newhistory []domain.Presensi
	for rows.Next() {
		history := domain.Presensi{}
		err := rows.Scan(&history.IdPresensi, &history.IdUser, &history.TanggalPresensi, &history.JamMasuk, &history.JamPulang, &history.TanggalPulang, &history.KeteranganMasuk, &history.KeteranganKeluar, &history.Latitude, &history.Longitude, &history.Alamat, &history.Status)
		helper.PanicIfError(err)
		newhistory = append(newhistory, history)

	}
	var gagal error
	if newhistory == nil {
		gagal = errors.New("history not found")
	} else {
		gagal = nil
	}
	return newhistory, gagal
}
