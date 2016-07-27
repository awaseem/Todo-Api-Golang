package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/todo", listTodo).Methods("GET")
	r.HandleFunc("/todo", addTodo).Methods("POST")
	r.HandleFunc("/todo", editTodo).Methods("PUT")
	r.HandleFunc("/todo", deleteTodo).Methods("DELETE")
	log.Println("Running Server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("hello world"))
}

func listTodo(w http.ResponseWriter, r *http.Request) {
	j, err := json.Marshal(todoList)
	if err != nil {
		log.Panic(err)
	} else {
		w.Header().Set("Content=Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

func addTodo(w http.ResponseWriter, r *http.Request) {
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

func editTodo(w http.ResponseWriter, r *http.Request) {
	var editTodo todo
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &editTodo)
	w.Header().Set("Content=Type", "application/json")
	for i, e := range todoList {
		if editTodo.ID == e.ID {
			// edit a todo based on the index
			todoList[i] = editTodo
			b, _ := json.Marshal(todoList)
			w.WriteHeader(http.StatusOK)
			w.Write(b)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("{ \"Message\": \"Unable to find todo with that ID:" + strconv.Itoa(editTodo.ID) + "\"}"))
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	var deleteTodo todo
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &deleteTodo)
	w.Header().Set("Content=Type", "application/json")
	for i, e := range todoList {
		if deleteTodo.ID == e.ID {
			// delete a todo based on the index
			todoList = todoList[:i+copy(todoList[i:], todoList[i+1:])]
			b, _ := json.Marshal(todoList)
			w.WriteHeader(http.StatusOK)
			w.Write(b)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("{ \"Message\": \"Unable to find todo with that ID: " + strconv.Itoa(deleteTodo.ID) + "\"}"))
}
