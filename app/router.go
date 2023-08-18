package app

import (
	"golang/rest-api-presensi/controller"
	"golang/rest-api-presensi/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	UserController controller.UserController,
	OtpController controller.OtpController,
	OfficeController controller.OfficeController,
	PresensiController controller.PresensiController,
) *httprouter.Router {

	router := httprouter.New()

	router.GET("/api/karyawancheck/:userEmail/:userIdKaryawan", UserController.KarywanCheck)
	router.GET("/api/statuskaryawancheck/:userEmail", UserController.StatusKarywanCheck)
	router.GET("/api/usercheck/:userEmail/:userIdKaryawan", UserController.EmailCheck)
	router.PUT("/api/presensikeluar", PresensiController.PresensiKeluar)
	router.GET("/api/riwayatpresensi/:userId", PresensiController.Riwayat)
	router.GET("/api/presensicheck/:userId", PresensiController.PresensiCheck)
	router.POST("/api/presensimasuk", PresensiController.PresensiMasuk)
	router.POST("/api/presensitidakmasuk", PresensiController.PresensiTidakMasuk)
	router.POST("/api/sendotp", OtpController.SendOtp)
	router.POST("/api/otpvalidation", OtpController.VerifikasiOtp)
	router.POST("/api/register", UserController.CreateAkun)
	router.POST("/api/recognition", UserController.CreateRecognition)
	router.POST("/api/login", UserController.Login)
	router.PUT("/api/updatepassword", UserController.UpdatePassword)
	router.GET("/api/profile/:userId", UserController.GetProfile)
	router.GET("/api/recognition/:userId", UserController.GetRecognition)
	router.GET("/api/smartphonecheck/:userId", UserController.GetSmartphoneCheck)
	router.GET("/api/office/:officeId", OfficeController.GetOfficeData)
	router.GET("/api/getiduser/:userEmail", UserController.GetIdUser)
	router.POST("/api/updateava", UserController.UpdateAva)

	router.PanicHandler = exception.ErrorHandler

	return router

}
