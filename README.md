# Learn Go Lang

This sample test API build with go language 1.10

## Getting Started

* copy .env.example to .env
* give `MONGODB_CONN_STRING` on .env file a connection value

## How to run

### With Native Go

* Install Depedencies
> `go get`

* Run Package
> `go run ./*.go`


## Request Example

```
curl --request GET \
  --url http://localhost:3100/users \
```