package controller

import (
	"golang/rest-api-presensi/entity/web"
	"golang/rest-api-presensi/helper"
	"golang/rest-api-presensi/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserControllerImpl(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) CreateAkun(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	akunCreateRequest := web.CreateAkunRequest{}
	helper.ReadFromRequestBody(request, &akunCreateRequest)

	akunResponse := controller.UserService.CreateAkun(request.Context(), akunCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   akunResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) CreateRecognition(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	recognitionCreateRequest := web.CreateRecognitionRequest{}
	helper.ReadFromRequestBody(request, &recognitionCreateRequest)

	recognitionResponse := controller.UserService.CreateRecognition(request.Context(), recognitionCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   recognitionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loginRequest := web.LoginRequest{}
	helper.ReadFromRequestBody(request, &loginRequest)

	loginResponse := controller.UserService.Login(request.Context(), loginRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   loginResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) GetProfile(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	profileResponse := controller.UserService.GetProfile(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   profileResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) GetRecognition(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	recognitionResponse := controller.UserService.GetRecognition(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   recognitionResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) GetSmartphoneCheck(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	smartphoneCheckResponse := controller.UserService.GetSmartphoneCheck(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   smartphoneCheckResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) EmailCheck(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userEmail := params.ByName("userEmail")
	userIdKaryawan := params.ByName("userIdKaryawan")

	profileResponse := controller.UserService.EmailCheck(request.Context(), userEmail, userIdKaryawan)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   profileResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) UpdatePassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	updatePasswordRequest := web.UpdatePasswordRequest{}
	helper.ReadFromRequestBody(request, &updatePasswordRequest)

	updatePasswordResponse := controller.UserService.UpdatePassword(request.Context(), updatePasswordRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   updatePasswordResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) KarywanCheck(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userEmail := params.ByName("userEmail")
	userIdKaryawan := params.ByName("userIdKaryawan")

	karyawanResponse := controller.UserService.KaryawanCheck(request.Context(), userEmail, userIdKaryawan)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   karyawanResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) StatusKarywanCheck(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userEmail := params.ByName("userEmail")

	StatuskaryawanResponse := controller.UserService.StatusKaryawanCheck(request.Context(), userEmail)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   StatuskaryawanResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) GetIdUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userEmail := params.ByName("userEmail")

	getidResponse := controller.UserService.GetIdUser(request.Context(), userEmail)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   getidResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) UpdateAva(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	updateAvaRequest := web.UpdateAvaRequest{}
	helper.ReadFromRequestBody(request, &updateAvaRequest)

	updateAvaResponse := controller.UserService.UpdateAva(request.Context(), updateAvaRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   updateAvaResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
