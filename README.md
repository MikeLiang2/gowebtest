# Go Web Test: Gin vs Echo vs Beego

This project demonstrates how to use different Go web frameworks — **Gin**, **Echo**, and **Beego** — to build a simple message-receiving server.

Each server:
- Listens on port `8080`
- Accepts a `GET` request with a `msg` query parameter
- Logs the message on the terminal
- Responds to the client with a confirmation message


## Directory Structure

gowebtest/
├── go.mod
├── server/
│   ├── main.go              # Entry point — choose which framework to run
│   ├── gin_server.go        # Gin example
│   ├── echo_server.go       # Echo example
│   └── beego_server.go      # Beego example
└── client/
    └── main.go              # Sends a message to the server



## How to Run

### 1. Install dependencies

go mod tidy


### 2. Run a server (choose only one at a time)

#### Option A: Run entire folder (recommended)
cd server
go run .


Then, open `main.go` and uncomment the line corresponding to the server you want to run:

func main() {
    startGinServer()
    // startEchoServer()
    // startBeegoServer()
}


#### Option B: Run individual file directly
go run gin_server.go         # or echo_server.go / beego_server.go

> Only one `main()` should be active at a time.


### 3. Test the server

In a browser or terminal, visit:
http://localhost:8080/?msg=HelloFromBrowser


Or run the client:

cd client
go run main.go


Client output example:

Server replied: Hello Client, I got: HelloFromClient


Server terminal will log:

Received: HelloFromClient



## Framework Comparison Summary

| Framework | Style           | Advantages                     | Use Cases               |
|-----------|------------------|----------------------------------|--------------------------|
| **Gin**   | Functional, Fast | Simple syntax, great performance | REST APIs, microservices |
| **Echo**  | Functional, Rich | Middleware, HTTP/2, WebSocket   | Secure APIs, web apps    |
| **Beego** | OOP (MVC)        | Full-stack, ORM, logs, router    | Enterprise apps          |


## Dependencies

- [Gin](https://github.com/gin-gonic/gin)
- [Echo](https://github.com/labstack/echo)
- [Beego](https://github.com/beego/beego)


## Author

Mike Liang | [github.com/MikeLiang2](https://github.com/MikeLiang2)
