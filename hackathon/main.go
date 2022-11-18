package main

import (
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// /transactions で -取引の全履歴をGET -取引を作成/削除/編集
	http.HandleFunc("/transactions", Transactions)

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
