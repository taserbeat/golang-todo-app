package main // import "github.io/taserbeat/golang-todo-app"

import (
	"fmt"

	"github.io/taserbeat/golang-todo-app/modules/models"
)

func main() {
	fmt.Println(models.Db)

	/*
		fmt.Println(setting.Config.Port)
		fmt.Println(setting.Config.SQLDriver)
		fmt.Println(setting.Config.DbName)
		fmt.Println(setting.Config.LogFile)
	*/

	/*
		  user, _ := models.GetUser(2)
			user.CreateTodo("First Todo")
	*/

	t, _ := models.GetTodo(3)
	t.DeleteTodo()

}
