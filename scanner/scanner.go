package scanner

import (
	"fmt"
	"net/http"
)

//Scan functionality
func Scan(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Welcome to the Scanner!")
	fmt.Println("Endpoint Hit: Scanner")
}
