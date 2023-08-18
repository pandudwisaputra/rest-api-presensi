package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang/rest-api-presensi/entity/domain"
	"golang/rest-api-presensi/helper"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) CreateAkun(ctx context.Context, tx *sql.Tx, akun domain.User) domain.User {
	script := "insert into user(id_karyawan, email, ava, password, android_id) values (?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, script, akun.IdKaryawan, akun.Email, akun.Avatar, akun.Password, akun.AndroidId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	akun.IdUser = int(id)

	return akun
}

func (repository *UserRepositoryImpl) CreateRecognition(ctx context.Context, tx *sql.Tx, recognition domain.Recognition) domain.Recognition {
	script := "insert into recognition(id_user, `key`, name, location_left, location_top, location_right, location_bottom, embeddings, distance) values (?,?,?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, script, recognition.IdUser, recognition.Key, recognition.Name, recognition.LocationLeft, recognition.LocationTop, recognition.LocationRight, recognition.LocationBottom, recognition.Embeddings, recognition.Distance)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	recognition.IdRecognition = int(id)

	return recognition
}

func (repository *UserRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, akun domain.User) (domain.User, error) {
	script := "select id_user, email, password from user where email = ? limit 1"
	rows, err := tx.QueryContext(ctx, script, akun.Email)
	helper.PanicIfError(err)
	user := domain.User{}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.IdUser, &user.Email, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("Email Atau Password Salah")
	}
}



func (repository *UserRepositoryImpl) GetProfile(ctx context.Context, tx *sql.Tx, user int) (domain.User, error) {
	script := "SELECT user.id_user, user.id_karyawan, user.email, user.password, user.ava, karyawan.nama_lengkap, karyawan.no_hp FROM user JOIN karyawan ON user.id_karyawan = karyawan.id_karyawan WHERE user.id_user = ?"
	rows, err := tx.QueryContext(ctx, script, user)
	helper.PanicIfError(err)
	akun := domain.User{}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&akun.IdUser, &akun.IdKaryawan, &akun.Email, &akun.Password, &akun.Avatar, &akun.NamaLengkap, &akun.NoHp)
		helper.PanicIfError(err)
		return akun, nil
	} else {
		return akun, errors.New("User Tidak Ditemukan")
	}
}

func (repository *UserRepositoryImpl) GetRecognition(ctx context.Context, tx *sql.Tx, user int) (domain.Recognition, error) {
	script := "SELECT id_recognition, id_user, `key`, name, location_left, location_top, location_right, location_bottom, embeddings, distance FROM recognition WHERE id_user = ?"
	rows, err := tx.QueryContext(ctx, script, user)
	helper.PanicIfError(err)
	recognition := domain.Recognition{}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&recognition.IdRecognition, &recognition.IdUser, &recognition.Key, &recognition.Name, &recognition.LocationLeft, &recognition.LocationTop, &recognition.LocationRight, &recognition.LocationBottom, &recognition.Embeddings, &recognition.Distance)
		helper.PanicIfError(err)
		return recognition, nil
	} else {
		return recognition, errors.New("Data Recognition Tidak Ditemukan")
	}
}

func (repository *UserRepositoryImpl) GetSmartphoneCheck(ctx context.Context, tx *sql.Tx, user int) (domain.User, error) {
	script := "SELECT id_user, android_id FROM user WHERE id_user = ?"
	rows, err := tx.QueryContext(ctx, script, user)
	helper.PanicIfError(err)
	smartphoneCheck := domain.User{}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&smartphoneCheck.IdUser, &smartphoneCheck.AndroidId)
		helper.PanicIfError(err)
		return smartphoneCheck, nil
	} else {
		return smartphoneCheck, errors.New("Data User Tidak Ditemukan")
	}
}

func (repository *UserRepositoryImpl) EmailCheck(ctx context.Context, tx *sql.Tx, email, idKaryawan string) (string, error) {
	script := "select email, id_karyawan from user where email = ? and id_karyawan = ?"
	rows, err := tx.QueryContext(ctx, script, email, idKaryawan)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		return "", errors.New("Email Dan Kode Pegawai Sudah Diaktivasi")
	} else {
		return "Email dan Kode Pegawai Tersedia", nil
	}
}


func (repository *UserRepositoryImpl) UpdatePassword(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	script := "update user set password = ? where id_user = ? and email = ?"
	_, err := tx.ExecContext(ctx, script, user.Password, user.IdUser, user.Email)
	helper.PanicIfError(err)
	return user
}

func (repository *UserRepositoryImpl) KaryawanCheck(ctx context.Context, tx *sql.Tx, karyawan, idKaryawan string) (domain.Karyawan, error) {
	script := "SELECT karyawan.id_karyawan, karyawan.id_jabatan, karyawan.nama_lengkap, karyawan.alamat, karyawan.agama, karyawan.email, karyawan.no_hp, karyawan.pendidikan, jabatan.jabatan FROM karyawan JOIN jabatan ON karyawan.id_jabatan = jabatan.id_jabatan WHERE karyawan.email = ? AND karyawan.id_karyawan = ?"
	rows, err := tx.QueryContext(ctx, script, karyawan, idKaryawan)
	helper.PanicIfError(err)
	defer rows.Close()

	karyawanstruct := domain.Karyawan{}
	if rows.Next() {
		err := rows.Scan(&karyawanstruct.IdKaryawan, &karyawanstruct.IdJabatan, &karyawanstruct.NamaLengkap, &karyawanstruct.Alamat, &karyawanstruct.Agama, &karyawanstruct.Email, &karyawanstruct.NoHp, &karyawanstruct.Pendidikan, &karyawanstruct.Jabatan)
		helper.PanicIfError(err)
		return karyawanstruct, nil
	}
	return karyawanstruct, errors.New("Email / Kode Pegawai Tidak Terdaftar")
}

func (repository *UserRepositoryImpl) StatusKaryawanCheck(ctx context.Context, tx *sql.Tx, karyawan string) (domain.Karyawan, error) {
	script := "SELECT karyawan.id_karyawan, karyawan.id_jabatan, karyawan.nama_lengkap, karyawan.alamat, karyawan.agama, karyawan.email, karyawan.no_hp, karyawan.pendidikan, karyawan.status_karyawan, jabatan.jabatan FROM karyawan JOIN jabatan ON karyawan.id_jabatan = jabatan.id_jabatan WHERE karyawan.email = ?"
	rows, err := tx.QueryContext(ctx, script, karyawan)
	helper.PanicIfError(err)
	defer rows.Close()

	karyawanstruct := domain.Karyawan{}
	if rows.Next() {
		err := rows.Scan(&karyawanstruct.IdKaryawan, &karyawanstruct.IdJabatan, &karyawanstruct.NamaLengkap, &karyawanstruct.Alamat, &karyawanstruct.Agama, &karyawanstruct.Email, &karyawanstruct.NoHp, &karyawanstruct.Pendidikan, &karyawanstruct.StatusKaryawan, &karyawanstruct.Jabatan)
		helper.PanicIfError(err)
		return karyawanstruct, nil
	}
	return karyawanstruct, errors.New("Email / Kode Pegawai Tidak Terdaftar")
}

func (repository *UserRepositoryImpl) GetIdUser(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	script := "select id_user, email from user where email = ?"
	rows, err := tx.QueryContext(ctx, script, email)
	helper.PanicIfError(err)
	defer rows.Close()

	result := domain.User{}
	if rows.Next() {
		err := rows.Scan(&result.IdUser, &result.Email)
		helper.PanicIfError(err)
		return result, nil
	} else {
		return result, errors.New("Email Tidak Terdaftar")
	}
}

func (repository *UserRepositoryImpl) UpdateAva(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	script := "update user set ava = ? where id_user = ?"
	_, err := tx.ExecContext(ctx, script, user.Avatar, user.IdUser)
	helper.PanicIfError(err)
	return user
}
