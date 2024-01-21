package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Wash the clothes", Completed: false},
	{ID: "2", Item: "Eat the dinner", Completed: false},
	{ID: "3", Item: "Go to Nudge office", Completed: false},
	{ID: "4", Item: "Talk to Isha", Completed: false},
}

func getTodosById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("No todo found!")
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo
	err := context.BindJSON(newTodo)
	if err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodosById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "No such todo in the DB!"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodosById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "No such todo in the DB!"})
		return
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos/add", addTodo)
	router.Run("localhost:8080")
}
