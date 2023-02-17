package main

import (
	"github.com/devlulcas/todoom/cmd/lib/router"
	"github.com/devlulcas/todoom/cmd/todos/controller/http"
	"github.com/devlulcas/todoom/cmd/todos/repository"
	use_case "github.com/devlulcas/todoom/cmd/todos/use-case"
	"github.com/devlulcas/todoom/configs"
	"github.com/devlulcas/todoom/db"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	dbConfig := configs.GetDBConfig()
	apiConfig := configs.GetAPIConfig()

	conn, err := db.OpenConnection(dbConfig)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := db.CloseConnection(conn)
		if err != nil {
			panic(err)
		}
	}()

	router := router.NewRouter()

	todoRepo := repository.NewPostgresTodoRepository(conn)

	todoUseCase := use_case.NewTodoUseCase(todoRepo)

	todoHandler := http.NewTodoHandler(todoUseCase)

	router.GET("/todos", todoHandler.HandleGetById)
	router.GET("/todos/list", todoHandler.HandleGetAll)
	router.POST("/todos/create", todoHandler.HandleCreate)
	router.PUT("/todos/update", todoHandler.HandleUpdate)
	router.DELETE("/todos/remove", todoHandler.HandleDelete)

	router.ListenAndServe(apiConfig.Host, apiConfig.Port)
}
