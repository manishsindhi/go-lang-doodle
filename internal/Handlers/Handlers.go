package Handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Take Shower", Completed: false},
	{ID: "3", Item: "Finish Tutorial", Completed: false},
}

func getTodos(context *gin.Context) {

	context.IndentedJSON(http.StatusOK, todos)

}

func addTodos(context *gin.Context) {
	var requestTodo todo

	if err := context.BindJSON(&requestTodo); err != nil {
		return
	}

	todos = append(todos, requestTodo)

	context.IndentedJSON(http.StatusCreated, todos)
}

func getTodoById(id string) (*todo, error) {

	for i, t := range todos {

		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	fmt.Println(id)
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, "todo not found")
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodo(context *gin.Context) {
	id := context.Param("id")

	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, "todo not found")
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)

}

func startHandlers() {
	//fmt.Println("Starting Webserver")
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodo)
	router.POST("/todos/add", addTodos)

	//fmt.Println("Starting Webserver")
	router.Run("localhost:9090")
}
