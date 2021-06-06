package models

import (
	"log"
	"time"
)

// ユーザーモデル
type User struct {
	Id        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	UUID      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

// ユーザーを作成する
func (u *User) CreateUser() (err error) {
	cmd := `INSERT INTO users (
    uuid,
    name,
    email,
    password,
    created_at) values ($1, $2, $3, $4, $5)`

	_, err = Db.Exec(cmd, createUUID(), u.Name, u.Email, Encrypt(u.Password), time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

// ユーザーを取得する
func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `SELECT id, uuid, name, email, password, created_at FROM users WHERE id = $1`
	err = Db.QueryRow(cmd, id).Scan(
		&user.Id,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	return user, err
}

// ユーザーを更新する
func (u *User) UpdateUser() (err error) {
	cmd := `UPDATE users SET name = $1, email = $2 WHERE id = $3`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.Id)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// ユーザーを削除する
func (u *User) DeleteUser() (err error) {
	cmd := `DELETE FROM users WHERE id = $1`
	_, err = Db.Exec(cmd, u.Id)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// Eメールアドレスからユーザーを取得する
func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `SELECT id, uuid, name, email, password, created_at FROM users WHERE email = $1`
	err = Db.QueryRow(cmd, email).Scan(
		&user.Id, &user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt)

	return user, err
}

// セッションを作成し、作成したセッションを取得する
func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	createCmd := `INSERT INTO sessions (uuid, email, user_id, created_at) VALUES ($1, $2, $3, $4)`

	_, err = Db.Exec(createCmd, u.UUID, u.Email, u.Id, u.CreatedAt)
	if err != nil {
		log.Println(err)
	}

	selectCmd := `SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = $1 AND email = $2`
	err = Db.QueryRow(selectCmd, u.Id, u.Email).Scan(
		&session.Id,
		&session.UUID,
		&session.Email,
		&session.UserId,
		&session.CreatedAt,
	)

	return session, err
}

// 有効なセッションが存在するかをチェックする
func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1`
	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.Id,
		&sess.UUID,
		&sess.Email,
		&sess.UserId,
		&sess.CreatedAt,
	)

	if err != nil {
		valid = false
		return
	}

	if sess.Id != 0 {
		valid = true
	}
	return valid, err
}

// セッションを削除する
func (sess *Session) DeleteSessionByUUID() (err error) {
	cmd := `DELETE FROM sessions WHERE uuid = $1`
	_, err = Db.Exec(cmd, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}
