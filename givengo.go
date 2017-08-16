package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/player/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "You requested data on the player: , %q", html.EscapeString(r.URL.Path))
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}