# Go Assignment

## System Requirements
1. Git
2. Docker or Go

## How to Build and Run the App
1. Clone the project to your computer
```
git clone https://github.com/lichao0817/go_assignment.git
cd go_assignment
```

2. Build the app with docker:

  ```
  docker build -t app/main .
  docker run -p 8080:8080 -d app/main
  ```
  or you can run the application using go instead:
  ```
  go run main.go
  ```
## API Services
### URL
```
GET http://localhost:8080/users/{username}
```
### Example Request 1 (Valid Username)
```
curl -X GET 'http://localhost:8080/users/vgbootcamp'
```
### Example Response 1
``` json
{
   "data":{
      "username":"vgbootcamp",
      "views":51066404,
      "followers":320571,
      "language":"en",
      "game":"Super Smash Bros. Melee",
      "display_name":"VGBootCamp",
      "created_at":"2010-01-09T21:35:21Z",
      "bio":"Headed by VGBC | GimR, Video Game Boot Camp is the leading Live-Streamer and Content Creator for competitive Super Smash Bros. This includes Melee, Brawl, Smash WiiU, and 64! Learn.Play.Win! ",
      "is_streaming":false
   },
   "status":200
}
```
### Example Request 2 (Invalid Username)
```
curl -X GET 'http://localhost:8080/users/zgjlkqjtw'
```
### Example Response 2
``` json
{
   "error":"Bad Request",
   "status":400,
   "message":"Invalid username"
}
```
### Example Request 3 (Invalid API URL)
```
curl -X GET 'http://localhost:8080/hello'
```
### Example Response 3
``` json
{
   "error":"Invalid API",
   "status":400,
   "message":"The API requested is not currently supported please use http://localhost:8080/users/{username}"
}
```
