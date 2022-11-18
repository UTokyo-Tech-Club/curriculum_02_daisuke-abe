package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/oklog/ulid/v2"
)

func Transactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodGet:
		rows, err := db.Query("SELECT * FROM transaction")
		if err != nil {
			log.Printf("fail: db.Query, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		users := make([]TransactionGet, 0)
		for rows.Next() {
			var u TransactionGet
			if err := rows.Scan(&u.Id, &u.Fromwhom, &u.Towhom, &u.Message, &u.Point); err != nil {
				log.Printf("fail: rows.Scan, %v\n", err)

				if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
					log.Printf("fail: rows.Close(), %v\n", err)
				}
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			users = append(users, u)
		}

		// ②-4
		bytes, err := json.Marshal(users)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)

	case http.MethodPost:
		var u TransactionPost
		fmt.Println("got POST method")

		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			fmt.Println("Decode失敗")
			fmt.Printf("%+v\n", u)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Printf("%+v\n", u)

		ins, err := db.Prepare("INSERT INTO transaction VALUES(?, ?, ?, ?, ?)")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer ins.Close()
		fmt.Println("SQL prepared")

		id := ulid.Make()
		res, err := ins.Exec(id.String(), u.Fromwhom, u.Towhom, u.Message, u.Point)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Println("inserted to DB")

		lastInsertID, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		log.Println(lastInsertID)
		fmt.Println("id: " + id.String())

	case http.MethodDelete:
		var u TransactionDelete
		fmt.Println("got Delete method")

		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			fmt.Println("Decode失敗")
			fmt.Printf("%+v\n", u)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		ins, err := db.Prepare("DELETE FROM transaction WHERE id = (?);")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer ins.Close()
		fmt.Println("SQL prepared")

		if _, err := ins.Exec(u.Id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Println(u.Id + "Deleted from DB")

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
