package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.io/taserbeat/golang-todo-app/modules/setting"
)

var Db *sql.DB
var err error

const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)

func init() {
	Db, err = sql.Open(
		setting.Config.SQLDriver,
    // sslmode=[disable | require]
		fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", setting.Config.DbHost, setting.Config.DbPort, setting.Config.DbUser, setting.Config.DbPassword, setting.Config.DbName))
	if err != nil {
		log.Fatalln(err)
	}
	// https://github.com/blobmon/communicate/blob/master/sql/create_table_query.sql
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
    id SERIAL PRIMARY KEY,
    uuid text NOT NULL,
    name text,
    email text,
    password text,
    created_at timestamp,
    UNIQUE(uuid)
  )`, tableNameUser)

	Db.Exec(cmdU)

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
    id SERIAL PRIMARY KEY,
    content text,
    user_id INTEGER,
    created_at timestamp
  )`, tableNameTodo)

	Db.Exec(cmdT)

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
    id SERIAL PRIMARY KEY,
    uuid text NOT NULL,
    email text,
    user_id INTEGER,
    created_at timestamp,
    UNIQUE(uuid)
  )`, tableNameSession)

	Db.Exec(cmdS)

}

func createUUID() uuid.UUID {
	uuidObj, _ := uuid.NewUUID()
	return uuidObj
}

func Encrypt(rawPassword string) (cryptedPassword string) {
	cryptedPassword = fmt.Sprintf("%x", sha1.Sum([]byte(rawPassword)))
	return cryptedPassword
}
