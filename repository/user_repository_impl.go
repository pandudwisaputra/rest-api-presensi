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
	script := "insert into user(id_karyawan, email, nama_lengkap,ava , no_hp, password, status_login) values (?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, script, akun.IdKaryawan, akun.Email, akun.Nama, akun.Avatar, akun.NoHp, akun.Password, akun.StatusLogin)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	akun.IdUser = int(id)

	return akun
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
	script := "select email, nama_lengkap, no_hp, password, status_login, ava from user where id_user = ?"
	rows, err := tx.QueryContext(ctx, script, user)
	helper.PanicIfError(err)
	akun := domain.User{}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&akun.Email, &akun.Nama, &akun.NoHp, &akun.Password, &akun.StatusLogin, &akun.Avatar)
		helper.PanicIfError(err)
		return akun, nil
	} else {
		return akun, errors.New("User Tidak Ditemukan")
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
	script := "select id_karyawan, id_jabatan, nama_lengkap, foto, alamat, agama, email, no_hp, pendidikan from karyawan where email = ? and id_karyawan = ?"
	rows, err := tx.QueryContext(ctx, script, karyawan, idKaryawan)
	helper.PanicIfError(err)
	defer rows.Close()

	karyawanstruct := domain.Karyawan{}
	if rows.Next() {
		err := rows.Scan(&karyawanstruct.IdKaryawan, &karyawanstruct.IdJabatan, &karyawanstruct.NamaLengkap, &karyawanstruct.Foto, &karyawanstruct.Alamat, &karyawanstruct.Agama, &karyawanstruct.Email, &karyawanstruct.NoHp, &karyawanstruct.Pendidikan)
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
