# xURL

![logo](f2e/src/assets/logo.png)

An easy URL shortener service.

## Tech Stack

#### Frontend
- react 19
- zustand 5
- node.js 22
- typescript 5.9
- tailwind 4
- vite 8

#### Backend

- go 1.26
- gin 1.12
- gorm 1.31
- traefik 2.8


## Start for Development

### Frontend

All action should under `f2e` folder.

Make sure you already under `node.js 22` and install all dependencies by `npm install`.

And then, you can run it.

    $ npm run dev

Frontend will start in `http://localhost:5173/`

### Backend

All action should under `apiserver` folder.

Make sure you already under `go 1.26` and install all dependencies by `go mod download`.

    $ go run main.go

It will start a web service in 8080 port.

You can test it from `http://localhost:8080/api/ping`. If you get `pong`, service would work fine.

### Test

Only backend now.

    $ go test -v ./tests


## Launch Services

Build images first, then you can launch them.

```bash
# build all images
$ docker-compose build

# launch all services
$ docker-compose up
```

Because of Traefik (reverse proxy), you can access the page by `http://xurl.localhose`.
