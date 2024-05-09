package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	auth "github.com/abbot/go-http-auth"
)

func Secret(user, realm string) string {
	if user == "wagner" {
		return "$1$edXBI6bl$VS.1bCYKWwvIIfbXNhYTn."
	}
	return ""
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Use: go run main.go <directory> <port>")
		os.Exit(1)
	}

	httpDir := os.Args[1]
	httpPort := os.Args[2]

	authenticator := auth.NewBasicAuthenticator("http-server", Secret)
	http.HandleFunc("/", authenticator.Wrap(func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
		http.FileServer(http.Dir(httpDir)).ServeHTTP(w, &r.Request)
	}))

	fmt.Printf("subindo servidor na porta %s", httpPort)
	log.Fatal(http.ListenAndServe(":"+httpPort, nil))
}
