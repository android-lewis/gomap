package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/android-lewis/gomap/scanner"
	"github.com/gin-gonic/gin"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	r := gin.Default()

	r.GET("/", scanner.Scan)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
