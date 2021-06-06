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
