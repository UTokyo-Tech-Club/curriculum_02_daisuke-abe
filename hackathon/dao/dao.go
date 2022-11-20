package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PWD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	_db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}
	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, %v\n", err)
	}
	Db = _db
}

// Ctrl+CでHTTPサーバー停止時にDBをクローズする
func CloseDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		if err := Db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}

func PointsCheck() (*sql.Rows, error) {
	return Db.Query("SELECT name, Sum(point) FROM user JOIN transaction ON transaction.towhom = user.id GROUP BY towhom ORDER BY Sum(point) DESC")
}

func TransactionsCheck() (*sql.Rows, error) {
	return Db.Query("SELECT * FROM transaction")
}

func TransactionCreate() (*sql.Stmt, error) {
	return Db.Prepare("INSERT INTO transaction VALUES(?, ?, ?, ?, ?)")
}

func TransactionUpdate() (*sql.Stmt, error) {
	return Db.Prepare("UPDATE transaction SET message = (?), point = (?) WHERE id = (?)")
}

func TransactionDelete() (*sql.Stmt, error) {
	return Db.Prepare("DELETE FROM transaction WHERE id = (?);")
}