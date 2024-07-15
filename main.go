package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func main() {
	fmt.Println("Hello World")

	now := time.Now().UTC()

	item := TodoItem{
		Id:          1,
		Title:       "Task 1",
		Description: "abc",
		Status:      "Doing",
		CreatedAt:   &now,
		UpdatedAt:   nil,
	}

	// from struct to json
	jsonData, err := json.Marshal(item)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))

	// from json to struct
	jsonStr := "{\"id\":1,\"title\":\"Task 1\",\"description\":\"abc\",\"status\":\"Doing\",\"created_at\":\"2024-07-15T16:25:41.2488272Z\",\"updated_at\":null}"

	var item2 TodoItem

	if err := json.Unmarshal([]byte(jsonStr), &item2); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(item2)
}
