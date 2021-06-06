package main // import "github.io/taserbeat/golang-todo-app"

import (
	"fmt"
	"log"

	"github.io/taserbeat/golang-todo-app/modules/setting"
)

func main() {
	fmt.Println(setting.Config.Port)
	fmt.Println(setting.Config.SQLDriver)
	fmt.Println(setting.Config.DbName)
	fmt.Println(setting.Config.LogFile)

	log.Println("test")
}
