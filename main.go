package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type todo struct {
	ID      int
	Done    bool
	Message string
}

type todos []todo

var todoList = todos{
	todo{ID: 1, Done: false, Message: "Get Eggs!"},
	todo{ID: 2, Done: false, Message: "Get Milk!"},
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/todo", ListTodos).Methods("GET")
	r.HandleFunc("/todo", AddTodo).Methods("POST")
	log.Println("Running Server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// HomeHandler provides route handling for the home route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("hello world"))
}

// ListTodos provides route handling for the list todos route
func ListTodos(w http.ResponseWriter, r *http.Request) {
	j, err := json.Marshal(todoList)
	if err != nil {
		log.Panic(err)
	} else {
		w.Header().Set("Content=Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

// AddTodo provides route handling that adds a todo
func AddTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo todo
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &newTodo)

	newTodo.ID = todoList[len(todoList)-1].ID + 1
	todoList = append(todoList, newTodo)

	j, _ := json.Marshal(todoList)
	w.Header().Set("Content=Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
