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

Forked and clone. cd inside the project.

`go run .`

###### Run test ?

testing state

`go test -v ./dto`

test route

`go test -v ./controller`

test services

`go test -v ./service`

###### With Docker ?



#### LICENSE
MIT
