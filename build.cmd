@echo off
setlocal

set arg=%~1

echo ** building frontend into web/app.wasm **
cmd /c build-web.cmd

if "%arg%" == "run" (
    go run .
) else (
    echo ** build backend **
    go build
    echo !! build fininshed !!
) 
