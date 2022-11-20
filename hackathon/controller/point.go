package controller

import (
	"UTTC_curriculum/test/usecase"

	"net/http"
)

func Point() {
	http.HandleFunc("/points", usecase.Points)
}
