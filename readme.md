## MSGO - Microservices in Go

### Hexagonal Architecture (Ports & Adapter)

[Wiki: Hexagonal Architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software))

###### Architecture benefits:
- loosely coupled
- interchangeable core application, UI, data objects, test
- testable
- flexible (e.g: can change database easily MySQL to MongoDB or any Database)

It might not be suitable for small application.

More information about this [here](https://dzone.com/articles/hexagonal-architecture-what-is-it-and-how-does-it)

###### How to run ?

Go Forked and clone this project and cd inside msgo.

`mv ./config/sample.conf.yaml ./config/conf.yaml`

`go run .`

###### Run test ?

testing state

`go test -v ./dto`

testing route

`go test -v ./controller`

testing services

`go test -v ./service`

test all with one liners

`go test --v ./service ./controller ./dto`

###### With Docker ?

Note: If you have mysql already running you can turn it off: `systemctl stop mysql`

`docker-compose up --build`

or

`docker-compose up`

#### LICENSE
MIT
