package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type todo struct {
	Done    bool
	Message string
}

type todos []todo

var todoList = todos{
	todo{Done: false, Message: "Get Eggs!"},
	todo{Done: false, Message: "Get Milk!"},
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/list", ListTodos)
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
	w.Header().Set("Content=Type", "application/json")
	w.WriteHeader(http.StatusOK)
	j, err := json.Marshal(todoList)
	if err != nil {
		log.Panic(err)
	} else {
		w.Write(j)
	}
}
