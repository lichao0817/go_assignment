package twitch

import (
  "net/http"
  "encoding/json"
  "time"
  "errors"
)

type ResponseData struct {
  Username string `json:"username"`
  Views int `json:"views"`
  Followers int `json:"followers"`
  Language string `json:"language"`
  Game string `json:"game"`
  DisplayNname string `json:"display_name"`
  CreatedAt time.Time `json:"created_at"`// users
  Bio string `json:"bio"`// users
  IsStreaming bool `json:"is_streaming"`// streams
}

type StreamInfo struct {
  Stream interface{}
}

type ChannelInfo struct {
  Views int
  Followers int
  Game string
  Language string
  Error string
}

type UserInfo struct {
  Display_name string // users
  Created_at time.Time // users
  Bio string // users
  Status int
  Error string
}

const clientId = "2vcql4x9t2rv57oadqc3uvbrzt44k4"
const apiUrl = "https://api.twitch.tv/kraken/"

func getActive(login string) (bool, error) {
  resp, err := getResponse("streams", login)
  if err != nil {
    return false, err
  }

  decoder := json.NewDecoder(resp.Body)
  var data StreamInfo
  err = decoder.Decode(&data)
  if err != nil {
    return false, err
  }
  return data.Stream != nil, nil
}

func getChannelInfo(login string) (*ChannelInfo, error) {
  resp, err := getResponse("channels", login)
  if err != nil {
    return nil, err
  }

  decoder := json.NewDecoder(resp.Body)
  var data ChannelInfo
  err = decoder.Decode(&data)
  if err != nil {
    return nil, err
  }
  return &data, nil
}

func getUserInfo(login string) (*UserInfo, error) {
  resp, err := getResponse("users", login)
  if err != nil {
    return nil, err
  }

  decoder := json.NewDecoder(resp.Body)
  var data UserInfo
  err = decoder.Decode(&data)
  if err != nil {
    return nil, err
  }
  if data.Error != "" {
    return nil, errors.New("invalid username")
  }
  return &data, nil
}

func newGetRequest(service string, login string) (*http.Request, error) {
  var req, err = http.NewRequest("GET", apiUrl + service + "/" + login, nil)
  if err != nil {
    return nil, err
  }
  req.Header.Add("Client-ID", clientId)
  return req, nil
}

func sendRequest(req *http.Request) (*http.Response, error) {
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return nil, err
  }
  return resp, nil
}

func getResponse(service string, login string) (*http.Response, error) {
  req, err := newGetRequest(service, login)
  if err != nil {
    return nil, err
  }
  return sendRequest(req)
}

func GetUserData(login string) (*ResponseData, error) {
  isActive, err := getActive(login)
  if err != nil {
    return nil, err
  }
  channelInfo, err := getChannelInfo(login)
  if err != nil {
    return nil, err
  }
  userInfo, err := getUserInfo(login)
  if err != nil {
    return nil, err
  }

  responseData := ResponseData{ login, channelInfo.Views, channelInfo.Followers,
    channelInfo.Language, channelInfo.Game, userInfo.Display_name, userInfo.Created_at,
  userInfo.Bio, isActive }

  return &responseData, nil
}
