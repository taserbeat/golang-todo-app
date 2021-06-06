package main // import "github.io/taserbeat/golang-todo-app"

import (
	"fmt"

	"github.io/taserbeat/golang-todo-app/modules/controllers"
	"github.io/taserbeat/golang-todo-app/modules/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

}
