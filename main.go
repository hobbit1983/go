package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// handleRequest handles incoming HTTP requests
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	fmt.Println("Request Inbound")
	// var jsonData map[string]interface{}

	fmt.Println(r.Header)

	b, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(b))
	// err := json.NewDecoder(r.Body).Decode(&jsonData)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// // Print the JSON data to the terminal
	// fmt.Printf("Received JSON: %+v\n", r.)

	// // Send a response back to the client
	w.WriteHeader(http.StatusOK)
}

func main() {
	// Set up the HTTP server
	http.HandleFunc("/", handleRequest)

	// Start the server
	fmt.Println("Server is listening on port 8573")
	log.Fatal(http.ListenAndServe(":8573", nil))
}





