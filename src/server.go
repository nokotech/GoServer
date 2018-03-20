package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"./common"
)

// ===============================================================
// データ構造
// ===============================================================
type Page struct {
	Title string
	Count int
}

// ===============================================================
//
// ===============================================================
func errorHandler(w http.ResponseWriter) {
	fmt.Fprint(w, "ERROR")
}

func handler1(w http.ResponseWriter, r *http.Request) {
	common.InfoLog("start....")
	common.InfoResLog(w, "")

	// テンプレート用のデータ
	page := Page{"Hello World.", 1}

	// JSONで返却
	json, _ := json.Marshal(page)
	fmt.Fprint(w, string(json))

	/*
		// ParseFilesを使う
		// tmpl, err := template.ParseFiles("./src/html/layout.html")
		// テンプレート文字列
		tmpl, err := template.New("new").Parse("{{.Title}} {{.Count}} count")
		if err != nil {
			return panic(err)
		}
	*/
	/*
		// テンプレートをジェネレート
		err = tmpl.Execute(w, page)
		if err != nil {
			return panic(err)
		}
	*/
}

// ===============================================================
// ハンドラを登録してウェブページを表示させる
// ===============================================================
func main() {
	common.InfoLog("start server.")
	http.HandleFunc("/", handler1)
	http.ListenAndServe(":3000", nil)
}
