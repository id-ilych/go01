# Go01 Project
This project is an exploration of the Go programming language. The aim is to gain experience in Go development and understand the language’s paradigms.

## Disclaimer
That is **_NOT_** an example of the **_PROPER_** way to write a web server or to interact with Kafka and a database.

## Project Structure
The project consists of two executables:

* `cmd/web` listens for HTTP POST requests containing `Hotdog` objects and produces Kafka messages accordingly.
* `cmd/consumer` listens for Kafka messages and stores the received `Hotdog` objects in a database.

## Getting Started
* `docker compose up -d` to spin up a local Kafka broker
* Start 3 terminals
    * Start the web server:
      `go run cmd/web/main.go`
    * Start the Kafka consumer:
      `go run cmd/consumer/main.go`
    * Send a message:
      `curl -X POST -H "Content-Type: application/json" -d '{"Title":"Classic Hotdog", "Calories":250, "Price":5}' http://localhost:8080/hotdogs`
* In the corresponding terminals you can see logs of what is going on in each of the processes

## Conclusions so far
Generally, it's ok. The code is readable though a little verbose to my taste especially when it comes to:
* explicit error handling even though in many cases the error is just returned up the call hierarchy and not _actually_ handled - some shortcut syntax for that typical case could've helped I think
* explicit `defer`s to "close" the resource, especially when the function may return an error and you have to wrap the call in a closure to log it

On the other hand, that explicitness might be beneficial in the long run when the code is read and modified by other people (you-from-tomorrow is another person too). But I don't have enough industrial experience with Go to be sure.

LLMs are very helpful in writing Go code.

So far I have no idea how this code should be covered with tests, though.

## Notable contributors
* [id-ilych](http://github.com/id-ilych)
* [Bing Chat](https://www.bing.com/search?q=Bing+AI&showconv=1)