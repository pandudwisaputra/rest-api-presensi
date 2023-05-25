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
		Nama:        request.Nama,
		NoHp:        request.NoHp,
		Password:    pwhash,
		StatusLogin: "0",
		Avatar:      "-",
	}

	akun = service.UserRepository.CreateAkun(ctx, tx, akun)

	return helper.ToCreateAkunResponse(akun)
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
	login.StatusLogin = "success"

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
