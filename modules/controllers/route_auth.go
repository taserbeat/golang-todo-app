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
		_, err := session(w, r)
		if err != nil {
			// ログインしていない場合
			generateHtml(w, nil, "layout", "public_navbar", "signup")
		} else {
			// ログイン済みの場合
			http.Redirect(w, r, "/todos", 302)
		}
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

// loginのハンドラ
func login(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		// ログインしていない場合
		generateHtml(w, nil, "layout", "public_navbar", "login")
	} else {
		// ログイン済みの場合
		http.Redirect(w, r, "/todos", 302)
	}
}

// 認証ハンドラ
func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", 302)
	}

	// パスワードが異なる場合
	if user.Password != models.Encrypt(r.PostFormValue("password")) {
		// ログインページにリダイレクト
		http.Redirect(w, r, "/login", 302)
	}

	// 以後、パスワードが等しい場合

	// セッションを作成する
	session, err := user.CreateSession()
	if err != nil {
		log.Println(err)
	}

	// クッキーをセットしてトップページにリダイレクト
	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    session.UUID,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/top", 302)
}

// logoutハンドラ
func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}

	if err != http.ErrNoCookie {
		// クッキーが存在する場合
		session := models.Session{
			UUID: cookie.Value,
		}

		// DBのセッション情報を削除する
		if session.DeleteSessionByUUID() == nil {
			// セッション情報の削除に成功した(エラーが返ってこない)場合はクッキーを削除する
			cookie.MaxAge = -1
			http.SetCookie(w, cookie)
		}
	}
	http.Redirect(w, r, "/login", 302)
}
