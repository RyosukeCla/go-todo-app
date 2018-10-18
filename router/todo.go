package router

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/RyosukeCla/go-todo-app/db"
	"github.com/go-chi/chi"
	"github.com/rs/xid"
)

type Todo struct {
	Id        string    `json:"id"`
	Text      string    `json:"text"`
	Checked   bool      `json:"checked"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type AddRequest struct {
	Text string `json:"text"`
}

type CheckRequest struct {
	Id string `json:"id"`
}

type DoneRequest struct {
	Ids []string `json:"ids"`
}

func doneTodo(w http.ResponseWriter, r *http.Request) {
	var doneRequest DoneRequest
	if err := BindJsonBody(r.Body, &doneRequest); err != nil {
		panic(err)
	}

	client := db.GetClient()

	for _, id := range doneRequest.Ids {
		client.Delete("todos", id)
	}

	WriteJson(w, doneRequest)
}

func checkTodo(w http.ResponseWriter, r *http.Request) {
	var checkRequest CheckRequest
	if err := BindJsonBody(r.Body, &checkRequest); err != nil {
		panic(err)
	}

	client := db.GetClient()

	var todo Todo
	if err := client.Read("todos", checkRequest.Id, &todo); err != nil {
		panic(err)
	}

	todo.Checked = !todo.Checked
	todo.UpdatedAt = time.Now()

	if err := client.Write("todos", todo.Id, todo); err != nil {
		panic(err)
	}

	WriteJson(w, todo)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	var addRequest AddRequest
	if err := BindJsonBody(r.Body, &addRequest); err != nil {
		panic(err)
	}

	newTodo := Todo{
		Id:        xid.New().String(),
		Text:      addRequest.Text,
		Checked:   false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	client := db.GetClient()
	if err := client.Write("todos", newTodo.Id, newTodo); err != nil {
		panic(err)
	}

	WriteJson(w, newTodo)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	client := db.GetClient()

	rawTodos, _ := client.ReadAll("todos")
	todos := make([]Todo, len(rawTodos))
	for index, rawTodo := range rawTodos {
		json.Unmarshal([]byte(rawTodo), &todos[index])
	}

	WriteJson(w, todos)
}

func TodoRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/add", addTodo)
	r.Post("/check", checkTodo)
	r.Post("/done", doneTodo)
	r.Get("/get", getTodo)
	return r
}
