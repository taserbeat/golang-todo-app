package controllers

import (
	"log"
	"net/http"
	"sort"

	"github.io/taserbeat/golang-todo-app/modules/models"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHtml(w, "Hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		// セッション情報を取得できなかった場合
		http.Redirect(w, r, "/login", 302)
	} else {
		// セッション情報を取得できた場合
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}
		todos, _ := user.GetTodoByUser()
		sort.SliceStable(todos, func(i, j int) bool { return todos[i].CreatedAt.After(todos[j].CreatedAt) })
		user.Todos = todos

		generateHtml(w, user, "layout", "private_navbar", "index")
	}
}

func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHtml(w, nil, "layout", "private_navbar", "todo_new")
	}
}

func todoSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		// セッションがエラーになる場合はログインページにリダイレクト
		http.Redirect(w, r, "/login", 302)
	} else {
		// フォームリクエストのエラーチェック
		err = r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}

		// ユーザー情報を取得
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		// リクエストのの内容からTodoデータをDBに登録
		content := r.PostFormValue("content")
		if err := user.CreateTodo(content); err != nil {
			log.Println(err)
		}

		// Todo一覧ページにリダイレクト
		http.Redirect(w, r, "/todos", 302)

	}
}

func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todo, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}

		generateHtml(w, todo, "layout", "private_navbar", "todo_edit")
	}
}

func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		content := r.PostFormValue("content")
		todo := models.Todo{
			Id:      id,
			Content: content,
			UserId:  sess.UserId,
		}

		if err := todo.UpdateTodo(); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/todos", 302)
	}
}

func todoDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}

		todo, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}

		if err := todo.DeleteTodo(); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/todos", 302)
	}
}
