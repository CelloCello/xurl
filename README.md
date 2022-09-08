# xURL

An easy URL shortener service.

## Frontend
- react 18.2
- node.js 14.20.0
- typescript 4.6.4
- tailwind 3.1.8
- vite 3.0.7


## Backend

- go 1.18
- gin 1.8.1
- gorm 1.23.6

All action should under `apiserver` folder.

### Start

    $ go run main.go

It will start a web service in 8080 port.
You can test it from `http://localhost:8080/ping`. If you get `pong`, service would work fine.

### Test

    $ go test -v ./tests

