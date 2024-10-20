package main

import (
    "fmt"
    "net/http"
    "log"
//    "log/slog"
)

func hello(w http.ResponseWriter, req *http.Request) {
    log.Println("GET /hello")
    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
    log.Println("GET /headers")
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {
    log.SetFlags(log.LstdFlags | log.Lmicroseconds)

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

    http.ListenAndServe(":8090", nil)
}

