package controller

import (
	"UTTC_curriculum/test/usecase"

	"net/http"
)

func Transaction() {
	http.HandleFunc("/transactions", usecase.Transactions)
}
