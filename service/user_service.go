package service

import (
	"context"
	"golang/rest-api-presensi/entity/web"
)

type UserService interface {
	CreateAkun(ctx context.Context, request web.CreateAkunRequest) web.CreateAkunResponse
	CreateRecognition(ctx context.Context, request web.CreateRecognitionRequest) web.CreateRecognitionResponse
	Login(ctx context.Context, request web.LoginRequest) web.LoginResponse
	GetProfile(ctx context.Context, id_user int) web.GetProfileResponse
	GetRecognition(ctx context.Context, id_user int) web.GetRecognitionResponse
	GetSmartphoneCheck(ctx context.Context, id_user int) web.GetSmartphoneCheckResponse
	EmailCheck(ctx context.Context, email, idKaryawan string) web.GetEmailCheckResponse
	UpdatePassword(ctx context.Context, request web.UpdatePasswordRequest) web.UpdatePasswordResponse
	KaryawanCheck(ctx context.Context, karyawan, idKaryawan string) web.GetKaryawanResponse
	StatusKaryawanCheck(ctx context.Context, karyawan string) web.GetStatusKaryawanResponse
	GetIdUser(ctx context.Context, email string) web.GetIdUserResponse
	UpdateAva(ctx context.Context, request web.UpdateAvaRequest) web.UpdateAvaResponse
}
