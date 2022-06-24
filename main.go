package main

import (
	"fmt"
	"net/http"
	"time"
)

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong! %s", time.Now())
}

func main() {
	http.HandleFunc("/ping", pong)
	http.ListenAndServe(":3000", nil)
}
