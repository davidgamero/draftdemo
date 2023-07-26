package main

import (
    "fmt"
    "net/http"
)


// from https://yourbasic.org/golang/http-server-example/
func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Print(w, "Hello, from Draft!")
}
