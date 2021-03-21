package main

import (

	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)


func main() {
	
	port := ":4545"

	router := CreateAllRouter();
	
	fmt.Printf("Listen port %s\n", port)

	if err := http.ListenAndServe(port, handlers.CORS()(router)); err != nil {
		log.Fatal(err)
	}
}
