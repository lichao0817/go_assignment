package main

import (
  "fmt"
  "net/http"
  "log"
  "encoding/json"
  "./twitch"
)

type ErrorResponse struct {
  Error string `json:"error"`
  Status int `json:"status"`
  Message string `json:"message"`
}

type Response struct {
  Data *twitch.ResponseData `json:"data"`
  Status int `json:"status"`
}

func main() {
  http.HandleFunc("/users", userHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
  login := r.URL.Query().Get("login")
  fmt.Println("Getting data from twitch server for user " + login + "...")
  data, err := twitch.GetUserData(login)
  if err != nil {
    errorResponse := ErrorResponse{"Bad Request", 400, "Invalid username"}
    json.NewEncoder(w).Encode(errorResponse)
  } else {
    response := Response{data, 200}
    json.NewEncoder(w).Encode(response)
  }
}
