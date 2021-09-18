package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "My Awesome Go App")
}

func secondPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This webpage is amazing!!")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/second", secondPage)

}

func main() {
	fmt.Println("Go Web App Started on Port 3000")
	setupRoutes()
	http.ListenAndServe(":3000", nil)
}
