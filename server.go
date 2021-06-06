package main // import "github.io/taserbeat/golang-todo-app"

import (
	"fmt"

	"github.io/taserbeat/golang-todo-app/modules/models"
	"github.io/taserbeat/golang-todo-app/modules/setting"
)

func main() {
	fmt.Println(setting.Config.Port)
	fmt.Println(setting.Config.SQLDriver)
	fmt.Println(setting.Config.DbName)
	fmt.Println(setting.Config.LogFile)

	fmt.Println(models.Db)
}
