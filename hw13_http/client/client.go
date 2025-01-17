package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8080/v1/hello")
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	defer resp.Body.Close()
}
