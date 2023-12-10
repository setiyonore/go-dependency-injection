package main

import (
	_ "github.com/go-sql-driver/mysql"
	"go-restful-api/helper"
	"go-restful-api/middleware"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}
func main() {
	server := InitilizedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
