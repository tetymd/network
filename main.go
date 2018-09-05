package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.Handle("/", http.FileServer(http.Dir("./")))
    http.HandleFunc("/a", handler)
    http.ListenAndServe(":8000", nil)
}


func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, r.Method)
}
