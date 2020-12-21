package endpoints

import (
	"fmt"
	"http"
)

//Scan functionality
func Scan(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Welcome to the Scanner!")
	fmt.Println("Endpoint Hit: Scanner")
}
