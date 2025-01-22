package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8080/v1/getUser")
	if err != nil {
		fmt.Printf("ошибка запроса: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("ошибка HTTP-ответа: %d\n", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ошибка чтения: %v", err)
	}

	fmt.Println(string(body))

}
