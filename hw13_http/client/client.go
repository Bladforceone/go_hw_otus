package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Bladforceone/go_hw_otus/hw13_http/types"
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
		return
	}

	userGet := types.User{}

	json.Unmarshal(body, &userGet)

	fmt.Print(userGet)

	userPost := types.User{
		ID:   52,
		Name: "Валера",
		Age:  37,
	}

	resp, err = http.Post("http://127.0.0.1:8080/v1/getUser", "application/json", strings.NewReader(userPost.String()))
	if err != nil {
		fmt.Printf("ошибка в запросе:%v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("ошибка запроса: %v", err)
	}
}
