package controllers

import (
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Err string
}

type error interface {
	Error() string
}

func MagaAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("i am here")
}

func TestAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API live and kicking"))
}
