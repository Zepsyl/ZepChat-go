package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func chatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Read the message using the modern io.ReadAll function
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		// Print message to console
		fmt.Printf("[Message]: %s\n", string(body))
		fmt.Fprintln(w, "Message received!")
	} else {
		fmt.Fprintln(w, "Send a POST request with your message.")
	}
}

func main() {
	http.HandleFunc("/chat", chatHandler)
	port := "8080"
	log.Printf("Chat server starting on :%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
