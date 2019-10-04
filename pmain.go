package main

import (
	"fmt"
	"net/http"
	"time"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Println(t.Format("20060102150405"))
	fmt.Fprintf(w, "My Awesome khghvgkg hfghfgh Go App")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
}

func main() {
	fmt.Println("Go Web App Started on Port 3000")
	setupRoutes()
	http.ListenAndServe(":3000", nil)
}
