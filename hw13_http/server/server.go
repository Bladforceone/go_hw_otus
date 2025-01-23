package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Bladforceone/go_hw_otus/hw13_http/types"
)

func main() {
	ip := flag.String("ip", "127.0.0.1", "IP address")
	port := flag.String("port", "8080", "port")
	flag.Parse()

	println("сервер запущен")

	http.Handle("/v1/hello", loggingMiddleware(http.HandlerFunc(hello)))
	http.Handle("/v1/getUser", loggingMiddleware(http.HandlerFunc(getUser)))
	http.Handle("/v1/createUser", loggingMiddleware(http.HandlerFunc(createUser)))

	server := &http.Server{
		Addr:         *ip + ":" + *port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("%v", err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Запрос %s %s от %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "метод не разрешен"})
		return
	}

	log.Print("Отправление приветствия клиенту")
	fmt.Fprint(w, "Hello world!")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "метод не разрешен"})
		return
	}

	w.Header().Set("Content-Type", "application/json")

	user := types.CreateExampleUser()

	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("Ошибка при кодировании JSON: %v", err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"error": "ошибка сервера"})
		return
	}

	log.Printf("Отправлены данные пользователя: %+v", user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "метод не разрешен"})
		return
	}

	var newUser types.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "ошибка в декодировании JSON",
			"details": err.Error(),
		})
		return
	}

	if newUser.Name == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "имя пользователя не может быть пустым"})
		log.Print("Ошибка валидации: имя пользователя пустое")
		return
	}

	if newUser.Age <= 0 || newUser.Age > 120 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "возраст должен быть в диапазоне от 1 до 120"})
		log.Printf("Ошибка валидации: некорректный возраст пользователя %d", newUser.Age)
		return
	}

	log.Printf("Создан новый пользователь %+v\n", newUser)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "пользователь успешно создан",
		"user":    fmt.Sprintf("%+v", newUser),
	})
}
