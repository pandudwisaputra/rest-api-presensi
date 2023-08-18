package main

import (
	"golang/rest-api-presensi/helper"
	"golang/rest-api-presensi/middleware"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    ":4500",
		Handler: authMiddleware,
	}
}

func main() {
	go helper.AutoPresensi()
	server := InitializedServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
