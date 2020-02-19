package main

import (
	"fmt"
	"net/http"
        "os"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello Go v2, %s", os.Getenv("HOSTNAME"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

