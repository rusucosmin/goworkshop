package web

import (
	"net/http"
	"os"
	"fmt"
	"encoding/json"
	"github.com/rusucosmin/goworkshop/model"
	"github.com/gorilla/mux"
)

const API_PORT_NAME = "API_PORT"
const API_PORT_VALUE = "8000"

func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods("GET")
	router.HandleFunc("/books", getAllBooksHandler).Methods("GET")
	router.HandleFunc("/authors", getAllAuthorsHandler).Methods("GET")
	router.HandleFunc("/authors/{uuid}", getAuthorByUUIDHandler).Methods("GET")
	router.HandleFunc("/books/{uuid}", getBookByUUIDHandler).Methods("GET")
	var port = getPort()
	fmt.Println("+-------------------------------+")
	fmt.Printf("| Starting sever on port: %s\t|\n", port)
	fmt.Println("+-------------------------------+")
	http.Handle("/", router)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "{\"ping\":\"pong\"}")
}

func getAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	serializedContent, err := json.Marshal(model.Books)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, string(serializedContent))
}

func getAllAuthorsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	serializedContent, err := json.Marshal(model.Authors)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, string(serializedContent))
}

func getBookByUUIDHandler(w http.ResponseWriter, r *http.Request) {
	UUID := mux.Vars(r)["uuid"]
	w.Header().Set("Content-Type", "application/json")
	for _, book := range model.Books {
		if book.UUID == UUID {
			serializedContent, err := json.Marshal(book)
			if err != nil {
				panic(err)
			}
			fmt.Fprintln(w, string(serializedContent))
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func getAuthorByUUIDHandler(w http.ResponseWriter, r *http.Request) {
	UUID := mux.Vars(r)["uuid"]
	w.Header().Set("Content-Type", "application/json")
	for _, author := range model.Authors {
		if author.UUID == UUID {
			serializedContent, err := json.Marshal(author)
			if err != nil {
				panic(err)
			}
			fmt.Fprintln(w, string(serializedContent))
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func getPort() string {
	port := os.Getenv(API_PORT_NAME)
	if port != "" {
		return port
	} else {
		return API_PORT_VALUE
	}
}
