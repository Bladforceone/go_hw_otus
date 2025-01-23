package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := flag.String("url", "http://127.0.0.1:8080", "URL сервера")
	path := flag.String("path", "/v1/getUser", "путь до ресурса")
	method := flag.String("method", "GET", "HTTP-метод(GET или POST)")
	data := flag.String("data", "", "данные для POST-запроса в формате JSON")

	flag.Parse()

	fullpath := *url + *path

	switch strings.ToUpper(*method) {
	case "GET":
		sendGetRequest(fullpath)
	case "POST":
		sendPostRequest(fullpath, *data)
	default:
		fmt.Println("Клиент поддерживает только POST или GET методы")
		os.Exit(1)
	}
}

func sendGetRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Ошибка при выполнении Get-запроса: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	printJSONResponse(resp)
}

func sendPostRequest(url, data string) {
	resp, err := http.Post(url, "application/json", strings.NewReader(data))
	if err != nil {
		fmt.Printf("Ошибка при выполнении Post-запроса: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	printJSONResponse(resp)
}

func printJSONResponse(r *http.Response) {
	fmt.Printf("Статус ответа:%s\n", r.Status)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Ошибка при чтении тела ответа: %v", err)
		os.Exit(1)
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Printf("Ошибка при парсинге JSON: %v\n", err)
		fmt.Printf("Тело ответа(сырое):\n %s", string(body))
		return
	}

	fmt.Println("Тело ответа(JSON):")
	prettyJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Printf("Ошибка при форматировании JSON: %v\n", err)
		fmt.Println(string(body))
		return
	}

	fmt.Println(string(prettyJSON))
}
