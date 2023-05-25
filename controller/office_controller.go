package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type OfficeController interface {
	GetOfficeData(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
