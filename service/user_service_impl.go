package service

import (
	"context"
	"database/sql"
	"golang/rest-api-presensi/entity/domain"
	"golang/rest-api-presensi/entity/web"
	"golang/rest-api-presensi/exception"
	"golang/rest-api-presensi/helper"
	"golang/rest-api-presensi/repository"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate}
}

func (service *UserServiceImpl) CreateAkun(ctx context.Context, request web.CreateAkunRequest) web.CreateAkunResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	pwhash, err := helper.HashPassword(request.Password)
	helper.PanicIfError(err)

	akun := domain.User{
		Email:       request.Email,
		IdKaryawan:  request.IdKaryawan,
		Password:    pwhash,
		Avatar:      "-",
		AndroidId: request.AndroidId,
	}

	akun = service.UserRepository.CreateAkun(ctx, tx, akun)

	return helper.ToCreateAkunResponse(akun)
}

func (service *UserServiceImpl) CreateRecognition(ctx context.Context, request web.CreateRecognitionRequest) web.CreateRecognitionResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	recognition := domain.Recognition{
		IdUser: request.IdUser,
		Key: request.Key,
		Name: request.Name,
		LocationLeft: request.LocationLeft,
		LocationTop: request.LocationTop,
		LocationRight: request.LocationRight,
		LocationBottom: request.LocationBottom,
		Embeddings: request.Embeddings,
		Distance: request.Distance,
	}

	recognition = service.UserRepository.CreateRecognition(ctx, tx, recognition)

	return helper.ToCreateRecognitionResponse(recognition)
}

func (service *UserServiceImpl) Login(ctx context.Context, request web.LoginRequest) web.LoginResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	email := domain.User{
		Email: request.Email,
	}

	login, err := service.UserRepository.Login(ctx, tx, email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	_, err = helper.CheckPasswordHash(request.Password, login.Password, "email atau password salah")
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	login.Email = request.Email
	login.Status = "Success"

	return helper.ToLoginResponse(login)
}



func (service *UserServiceImpl) GetProfile(ctx context.Context, id_user int) web.GetProfileResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	profile, err := service.UserRepository.GetProfile(ctx, tx, id_user)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToGetProfileResponse(profile)
}

func (service *UserServiceImpl) GetRecognition(ctx context.Context, id_user int) web.GetRecognitionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	recognition, err := service.UserRepository.GetRecognition(ctx, tx, id_user)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToGetRecognitionResponse(recognition)
}

func (service *UserServiceImpl) GetSmartphoneCheck(ctx context.Context, id_user int) web.GetSmartphoneCheckResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	smartphoneCheck, err := service.UserRepository.GetSmartphoneCheck(ctx, tx, id_user)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToGetSmartphoneCheckResponse(smartphoneCheck)
}

func (service *UserServiceImpl) EmailCheck(ctx context.Context, email, idKaryawan string) web.GetEmailCheckResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	emailCheck, err := service.UserRepository.EmailCheck(ctx, tx, email, idKaryawan)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToGetEmailResponse(emailCheck)
}


func (service *UserServiceImpl) UpdatePassword(ctx context.Context, request web.UpdatePasswordRequest) web.UpdatePasswordResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	newPassword, err := helper.HashPassword(request.NewPassword)
	helper.PanicIfError(err)

	updatePassword := domain.User{
		IdUser:   request.IdUser,
		Email:    request.Email,
		Password: newPassword,
	}

	updatePassword = service.UserRepository.UpdatePassword(ctx, tx, updatePassword)

	return helper.ToUpdatePasswordResponse(updatePassword)
}

func (service *UserServiceImpl) KaryawanCheck(ctx context.Context, karyawan, idKaryawan string) web.GetKaryawanResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	karyawancheck, err := service.UserRepository.KaryawanCheck(ctx, tx, karyawan, idKaryawan)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToGetKaryawanResponse(karyawancheck)
}

func (service *UserServiceImpl) StatusKaryawanCheck(ctx context.Context, karyawan string) web.GetStatusKaryawanResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	statuskaryawancheck, err := service.UserRepository.StatusKaryawanCheck(ctx, tx, karyawan)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToGetStatusKaryawanResponse(statuskaryawancheck)
}

func (service *UserServiceImpl) GetIdUser(ctx context.Context, email string) web.GetIdUserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	idUserCheck, err := service.UserRepository.GetIdUser(ctx, tx, email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToGetIdUserResponse(idUserCheck)
}

func (service *UserServiceImpl) UpdateAva(ctx context.Context, request web.UpdateAvaRequest) web.UpdateAvaResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	updateAva := domain.User{
		IdUser: request.IdUser,
		Avatar: request.Ava,
	}

	updateAva = service.UserRepository.UpdateAva(ctx, tx, updateAva)
	return helper.ToUpdateAvaResponse(updateAva)
}
