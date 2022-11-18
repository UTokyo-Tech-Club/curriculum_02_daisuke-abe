package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/joho/godotenv"
	"github.com/oklog/ulid/v2"
)

type UserResForHTTPGet struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserResForHTTPPost struct {
	Name string
	Age  int
}

type TransactionPost struct {
	Fromwhom string
	Towhom   string
	Message  string
	Point    int
}

type TransactionGet struct {
	Id       string `json:"id"`
	Fromwhom string `json:"fromwhom"`
	Towhom   string `json:"towhom"`
	Message  string `json:"message"`
	Point    int    `json:"point"`
}

type TransactionDelete struct {
	Id string
}

type EditPost struct {
	Id      string
	Message string
	Point   int
}

type PointGet struct {
	Name  string `json:"name"`
	Point int    `json:"point"`
}

// ① GoプログラムからMySQLへ接続
var db *sql.DB

func init() {
	// ①-1
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PWD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	// ①-2
	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	_db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}
	// ①-3
	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, %v\n", err)
	}
	db = _db
}

func list(w http.ResponseWriter, r *http.Request) {
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
	}
}

// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodGet:
		// ②-1
		name := r.URL.Query().Get("name") // To be filled
		if name == "" {
			log.Println("fail: name is empty")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// ②-2
		rows, err := db.Query("SELECT id, name, age FROM user WHERE name = ?", name)
		if err != nil {
			log.Printf("fail: db.Query, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// ②-3
		users := make([]UserResForHTTPGet, 0)
		for rows.Next() {
			var u UserResForHTTPGet
			if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
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

	// case http.MethodPost:
	// 	var u TransactionPost
	// 	fmt.Println("got POST method")

	// 	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
	// 		fmt.Println("Decode失敗")
	// 		fmt.Printf("%+v\n", u)
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		return
	// 	}

	// 	fmt.Printf("%+v\n", u)

	// 	ins, err := db.Prepare("INSERT INTO transaction VALUES(?, ?, ?, ?, ?)")
	// 	if err != nil {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		return
	// 	}
	// 	defer ins.Close()
	// 	fmt.Println("SQL prepared")

	// 	id := ulid.Make()
	// 	res, err := ins.Exec(id.String(), u.Fromwhom, u.Towhom, u.Message, u.Point)
	// 	if err != nil {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		return
	// 	}
	// 	fmt.Println("inserted to DB")

	// 	lastInsertID, err := res.LastInsertId()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 	}
	// 	w.WriteHeader(http.StatusOK)
	// 	log.Println(lastInsertID)
	// 	fmt.Println("id: " + id.String())
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func edit(w http.ResponseWriter, r *http.Request) {
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
	}
}

func points(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodGet:
		rows, err := db.Query("SELECT name, Sum(point) FROM user JOIN transaction ON transaction.towhom = user.id GROUP BY towhom ORDER BY Sum(point) DESC")
		if err != nil {
			log.Printf("fail: db.Query, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		points := make([]PointGet, 0)
		for rows.Next() {
			var p PointGet
			if err := rows.Scan(&p.Name, &p.Point); err != nil {
				log.Printf("fail: rows.Scan, %v\n", err)

				if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
					log.Printf("fail: rows.Close(), %v\n", err)
				}
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			points = append(points, p)
		}

		// ②-4
		bytes, err := json.Marshal(points)
		if err != nil {
			log.Printf("fail: json.Marshal, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	}
}

func main() {
	// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
	// POSTもこっち
	http.HandleFunc("/transaction", handler)

	// /transactions で取引の全履歴をJsonで返す
	http.HandleFunc("/transactions", list)

	// 貢献の編集
	http.HandleFunc("/edit", edit)

	// /points で各ユーザーごとのポイント数を返す
	http.HandleFunc("/points", points)
	// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
	closeDBWithSysCall()

	// 8080番ポートでリクエストを待ち受ける
	log.Println("Listening...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
func closeDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}

//
