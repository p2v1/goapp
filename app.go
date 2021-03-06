package main

import (
    "fmt"
    "net/http"
    "os"
)

func hello(w http.ResponseWriter, req *http.Request) {
    username := os.Getenv("NAMEHOST")
    fmt.Fprintf(w, username)
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

    http.HandleFunc("/", hello)
    http.HandleFunc("/headers", headers)

    http.ListenAndServe(":8080", nil)
}
