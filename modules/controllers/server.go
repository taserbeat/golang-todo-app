package controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.io/taserbeat/golang-todo-app/modules/setting"
)

// 共通のレイアウトテンプレートを生成する
func generateHtml(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("modules/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

// サーバーを起動する
func StartMainServer() (err error) {
	// 静的ファイルのハンドラ
	files := http.FileServer(http.Dir(setting.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// ルートのハンドラ
	http.HandleFunc("/", top)

	// signupハンドラ
	http.HandleFunc("/signup", signup)
	return http.ListenAndServe(":"+setting.Config.Port, nil)
}
