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
  http.HandleFunc("/", defaultHandler)
  http.HandleFunc("/users/", userHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
func defaultHandler(w http.ResponseWriter, r *http.Request) {
  errorResponse := ErrorResponse{"Bad Request", 400, "The API requested is not currently supported please use http://localhost:8080/users/{username}"}
  json.NewEncoder(w).Encode(errorResponse)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
  login := r.URL.Path[len("/users/"):]
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
