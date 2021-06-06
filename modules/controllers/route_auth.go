package controllers

import (
	"log"
	"net/http"

	"github.io/taserbeat/golang-todo-app/modules/models"
)

// signupのハンドラ
func signup(w http.ResponseWriter, r *http.Request) {
	// GEtメソッドの時はサインアップページを返す
	if r.Method == "GET" {
		generateHtml(w, nil, "layout", "public_navbar", "signup")
	} else if r.Method == "POST" { // POSTメソッドの時はサインアップ処理を行う
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		// 入力フォームの値からユーザーモデルを作成し、データベースに登録する
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}

		// 成功したらトップページにリダイレクトする
		http.Redirect(w, r, "/", 303)
	}
}
