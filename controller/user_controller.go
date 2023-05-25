package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	CreateAkun(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetProfile(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	EmailCheck(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdatePassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	KarywanCheck(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetIdUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateAva(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
