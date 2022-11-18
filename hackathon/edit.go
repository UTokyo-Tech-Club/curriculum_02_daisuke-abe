package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodPost:
		var u EditPost
		fmt.Println("got Edit POST method")

		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			fmt.Println("Decode失敗")
			fmt.Printf("%+v\n", u)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Printf("%+v\n", u)

		ins, err := db.Prepare("UPDATE transaction SET message = (?), point = (?) WHERE id = (?)")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer ins.Close()
		fmt.Println("SQL prepared")

		if _, err := ins.Exec(u.Message, u.Point, u.Id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Println("DB edited")

		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
