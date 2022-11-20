package controller

import (
	"UTTC_curriculum/test/usecase"

	"net/http"
)

func Point() {
	http.HandleFunc("/transactions", usecase.Transactions)
}
