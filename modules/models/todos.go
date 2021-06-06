package models

import (
	"log"
	"time"
)

// Todoモデル
type Todo struct {
	Id        int
	Content   string
	UserId    int
	CreatedAt time.Time
}

// Todoを作成する
func (u *User) CreateTodo(content string) (err error) {
	cmd := `INSERT INTO todos (
    content,
    user_id,
    created_at) VALUES ($1, $2, $3)`

	_, err = Db.Exec(cmd, content, u.Id, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// Todoを1つ取得する
func GetTodo(id int) (todo Todo, err error) {
	cmd := `SELECT id, content, user_id, created_at FROM todos WHERE id = $1`
	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan(&todo.Id, &todo.Content, &todo.UserId, &todo.CreatedAt)
	if err != nil {
		log.Fatalln(err)
	}

	return todo, err
}

// Todoをすべて取得する
func GetTodos() (todos []Todo, err error) {
	cmd := "SELECT id, content, user_id, created_at FROM todos"
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.Id, &todo.Content, &todo.UserId, &todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

// 特定ユーザーのTodoをすべて取得する
func (u *User) GetTodoByUser() (todos []Todo, err error) {
	cmd := "SELECT id, content, user_id, created_at FROM todos WHERE user_id = $1"
	rows, err := Db.Query(cmd, u.Id)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.Id, &todo.Content, &todo.UserId, &todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

// Todoを更新する
func (t *Todo) UpdateTodo() (err error) {
	cmd := `UPDATE todos SET content = $1, user_id = $2 WHERE id = $3`
	_, err = Db.Exec(cmd, t.Content, t.UserId, t.Id)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

// Todoを削除する
func (t *Todo) DeleteTodo() (err error) {
	cmd := `DELETE FROM todos WHERE id = $1`
	_, err = Db.Exec(cmd, t.Id)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}
