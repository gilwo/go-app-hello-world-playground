@echo off
setlocal

set arg=%~1

echo ** building frontend into web/app.wasm **
cmd /c build-web.cmd

if "%arg%" == "run" (
    go run backend\main.go
) else (
    echo ** build backend **
    go build -o server.exe backend\main.go
    echo !! build fininshed !!
) 
