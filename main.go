package main

import (
	"fmt"
	"net/http"
	"github.com/pandikaadi/ai-chat-web/middleware"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World")

}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", helloHandler)
	fmt.Println("Server is starting on port 8080...")
	http.ListenAndServe {
		Addr: ":8080",
		Handler: middleware.Logging(router)
	}
}
