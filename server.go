package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
  "os"
)

// Ex: http://localhost:8000/?board=test
func handler(w http.ResponseWriter, r *http.Request) {
	board := r.URL.Query().Get("board")
	status := "fail"

	if validBoardLength(board) {

	}

	json.NewEncoder(w).Encode(map[string]string{"status": status, "board": board})
}

func main() {
	http.HandleFunc("/", handler)
  port := os.Getenv("PORT")
  if port == "" {
    port = "8000"
  }
  fmt.Printf("Server listening at port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
