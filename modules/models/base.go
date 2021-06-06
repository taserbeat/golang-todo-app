package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.io/taserbeat/golang-todo-app/modules/setting"
)

var Db *sql.DB
var err error

const (
	tableNameUser = "users"
)

func init() {
	Db, err = sql.Open(
		setting.Config.SQLDriver,
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

}
