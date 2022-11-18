package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
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

func main() {
	// ② /userでリクエストされたらnameパラメーターと一致する名前を持つレコードをJSON形式で返す
	// POSTもこっち
	// http.HandleFunc("/transaction", handler)

	// /transactions で取引の全履歴をJsonで返す
	http.HandleFunc("/transactions", Transactions)

	// 貢献の編集
	http.HandleFunc("/edit", Edit)

	// /points で各ユーザーごとのポイント数を返す
	http.HandleFunc("/points", Points)
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
