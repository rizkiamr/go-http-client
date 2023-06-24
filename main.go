package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var baseURL string = "https://jsonplaceholder.typicode.com"

func main() {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, baseURL+"/todos/1", nil)
	if err != nil {
		panic(err)
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexpected status: got %v", response.Status))
	}

	fmt.Println(response.Header.Get("Content-Type"))

	var data struct {
		UserID    int    `json:"userId"`
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", data)
}
