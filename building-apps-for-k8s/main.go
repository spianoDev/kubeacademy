package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    log.Println("Starting: Building Applications for K8s App")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

    log.Println("Request received from %s", r.RemoteAddr)
    fmt.Fprintf(w, "Hello from spianoDev's Building Apps for K8s")
}