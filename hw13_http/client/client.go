package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	url := flag.String("url", "http://127.0.0.1:8080", "URL сервера")
	path := flag.String("path", "/v1/getUser", "путь до ресурса")
	method := flag.String("method", "GET", "HTTP-метод(GET или POST)")
	data := flag.String("data", "", "данные для POST-запроса в формате JSON")

	flag.Parse()

	fullpath := *url + *path

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	switch strings.ToUpper(*method) {
	case "GET":
		sendGetRequest(ctx, fullpath)
	case "POST":
		sendPostRequest(ctx, fullpath, *data)
	default:
		fmt.Println("Клиент поддерживает только POST или GET методы")
		return
	}
}

func sendGetRequest(ctx context.Context, url string) {
	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Printf("Ошибка при создании GET-запроса: %v", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Ошибка при выполнении GET-запроса: %v", err)
		return
	}
	defer resp.Body.Close()

	printJSONResponse(resp)
}

func sendPostRequest(ctx context.Context, url, data string) {
	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(data))
	if err != nil {
		fmt.Printf("Ошибка при создании POST-запроса: %v", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Ошибка при выполнении POST-запроса: %v", err)
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
