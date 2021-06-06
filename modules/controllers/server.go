package controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.io/taserbeat/golang-todo-app/modules/models"
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

// リクエストからセッション情報を取得する
func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid Session")
		}
	}

	return sess, err
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

	// loginハンドラ
	http.HandleFunc("/login", login)

	// 認証ハンドラ
	http.HandleFunc("/authenticate", authenticate)

	// logoutハンドラ
	http.HandleFunc("/logout", logout)

	// indexハンドラ (要ログイン)
	http.HandleFunc("/todos", index)

	return http.ListenAndServe(":"+setting.Config.Port, nil)
}
