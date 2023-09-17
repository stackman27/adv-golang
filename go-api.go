package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2/log"
)

func api_main() {
	handlePostRequest()
}

// Define a struct that matches the JSON response structure
type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TodoPost struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"usedId"`
}

// Get Endpoint
func handleGetRequest() {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	// Send a GET request to the API
	resp, err := http.Get(url)
	if err != nil {
		log.Error("Error occoured while getting")
		return
	}

	// ensure that the response body of the HTTP request is closed properly when you're done processing the response.
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	// Parse the JSON response into a CatFact struct
	var todo Todo
	if err := json.Unmarshal(body, &todo); err != nil {
		fmt.Printf("Error parsing JSON response: %v\n", err)
		return
	}
	// Access the parsed data
	fmt.Println(todo)
}

// Post endpoint
func handlePostRequest() {
	// URL of the API you want to send the POST request to
	url := "https://jsonplaceholder.typicode.com/posts"

	// Create a test Todo instance
	testTodo := TodoPost{
		Title:  "sishir",
		Body:   "sishirRequest",
		UserId: 1,
	}

	// Serialize the testTodo struct into JSON
	requestBody, err := json.Marshal(testTodo)
	if err != nil {
		fmt.Printf("Error serializing JSON: %v\n", err)
		return
	}

	// Send the POST request
	response, err := http.Post(url, "application/json; charset=UTF-8", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Printf("Error sending POST request: %v\n", err)
		return
	}
	defer response.Body.Close()

	// Check the response status code
	if response.StatusCode != http.StatusCreated {
		fmt.Printf("Received non-OK status code: %v\n", response.StatusCode)
		return
	}

	// Read the response body
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	// Process the response body
	fmt.Println("Response Body:", string(responseBody))
}
