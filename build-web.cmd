@echo off
set GOARCH=wasm
set GOOS=js
go build -o web/app.wasm frontend\main.go