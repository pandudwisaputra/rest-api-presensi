package repository

import (
	"context"
	"database/sql"
	"golang/rest-api-presensi/entity/domain"
)

type UserRepository interface {
	CreateAkun(ctx context.Context, tx *sql.Tx, akun domain.User) domain.User
	CreateRecognition(ctx context.Context, tx *sql.Tx, recognition domain.Recognition) domain.Recognition
	Login(ctx context.Context, tx *sql.Tx, akun domain.User) (domain.User, error)
	GetProfile(ctx context.Context, tx *sql.Tx, user int) (domain.User, error)
	GetRecognition(ctx context.Context, tx *sql.Tx, user int) (domain.Recognition, error)
	GetSmartphoneCheck(ctx context.Context, tx *sql.Tx, user int) (domain.User, error)
	EmailCheck(ctx context.Context, tx *sql.Tx, email string, idKaryawan string) (string, error)
	UpdatePassword(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	KaryawanCheck(ctx context.Context, tx *sql.Tx, karyawan string, idKaryawan string) (domain.Karyawan, error)
	StatusKaryawanCheck(ctx context.Context, tx *sql.Tx, karyawan string) (domain.Karyawan, error)
	GetIdUser(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
	UpdateAva(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
}
