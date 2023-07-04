package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Example: go run main.go <name> <port>")
		os.Exit(1)
	}

	serverName := os.Args[1]
	serverPort := os.Args[2]

	fmt.Printf("starting server %s on port %s\n", serverName, serverPort)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("----> Received on server %s\n", serverName)

		fmt.Printf("URL: %s %s\n", r.Method, r.URL.String())

		fmt.Println("Headers:")
		for key, value := range r.Header {
			fmt.Printf("\t%s: %v\n", key, value)
		}

		w.Write([]byte("Hello From Server " + serverName))

		fmt.Println("<---- Response sent")
	})

	http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil)
}
