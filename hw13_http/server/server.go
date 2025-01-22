package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/Bladforceone/go_hw_otus/hw13_http/types"
)

func main() {
	ip := flag.String("ip", "127.0.0.1", "IP address")
	port := flag.String("port", "8080", "port")
	flag.Parse()

	println("run server")

	http.HandleFunc("/v1/hello", hello)
	http.HandleFunc("/v1/getUser", getUser)
	http.HandleFunc("/v1/createUser", createUser)

	if err := http.ListenAndServe(*ip+":"+*port, nil); err != nil {
		fmt.Printf("%v", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
	log.Printf("Было отправлено приветствие %s", r.URL.Host)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	user := types.User{
		ID:   1,
		Name: "Райн Гослинг",
		Age:  52,
	}

	json.NewEncoder(w).Encode(user)
	log.Printf("Отправлены данные пользователя: %+v", user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newUser types.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "error decoding JSON %w", err)
		return
	}

	log.Println("Создан новый пользователь")
	fmt.Printf("New user: %+v\n", newUser)

	w.WriteHeader(http.StatusCreated)
}
