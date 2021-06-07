package controllers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
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

var validPath = regexp.MustCompile("^/todos/(edit|update|delete)/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// /todos/edit/1
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, qi)
	}
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

	// 新規todo作成ページ用ハンドラ (要ログイン)
	http.HandleFunc("/todos/new", todoNew)

	// 新規todoの作成APIハンドラ (要ログイン)
	http.HandleFunc("/todos/save", todoSave)

	// todoの編集用ページ (要ログイン)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))

	// todoの更新APIのハンドラ (要ログイン)
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))

	// todoの削除APIのハンドラ (要ログイン)
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))

	return http.ListenAndServe(":"+setting.Config.Port, nil)
}
