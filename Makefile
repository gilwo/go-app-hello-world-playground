

all: build

run: run-server

build-web:
	GOARCH=wasm GOOS=js go build -o web/app.wasm frontend/main.go

build: build-web
	go build -o server backend/main.go

run-server:
	go run backend/main.go



