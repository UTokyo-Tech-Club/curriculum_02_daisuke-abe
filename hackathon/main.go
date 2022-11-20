package main

import (
	"UTTC_curriculum/test/controller"
	"UTTC_curriculum/test/dao"

	"log"
	"net/http"
)

func main() {

	// /transactions で -取引の全履歴をGET -取引を作成/削除/編集
	controller.Point()

	// /points で各ユーザーごとのポイント数を返す
	controller.Transaction()

	// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする
	dao.CloseDBWithSysCall()

	// 8080番ポートでリクエストを待ち受ける
	log.Println("Listening...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
