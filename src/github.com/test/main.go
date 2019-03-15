package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Load the index.html template.
var tmpl = template.Must(template.New("tmpl").ParseFiles("index.html"))

func main() {
	// Serve / with the index.html file.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.ExecuteTemplate(w, "index.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Serve /callme with a text response.
	http.HandleFunc("/callme", func(w http.ResponseWriter, r *http.Request) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintln(w, text)
	})

	// Start the server at http://localhost:9000
	fmt.Println("Server is currently up and running....")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
