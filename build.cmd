@echo off
setlocal


set argC=0
for %%x in (%*) do Set /A argC+=1

set arg1=%~1

echo %arg1%
if "%arg1%" == "-h" set USAGE=1
if "%arg1%" == "--help" set USAGE=1
if defined USAGE (
    goto :usage
)

if "%argC%" == "0" (
    set WEB=1
    set BUILD=1
    echo *** building frontend and backend ***
)
if "%arg1%" == "runall" (
    set WEB=1
    set RUN=1
    echo *** build and run frontend and backend ***
)
if "%arg1%" == "run" (
    set RUN=1
    echo *** build and run backend ***
)
if "%arg1%" == "web" (
    set WEB=1
    echo *** build frontend ***
)

if defined WEB (
echo ** building frontend into web/app.wasm **
cmd /c build-web.cmd
)

if defined RUN (
    go run backend\main.go
) else if defined BUILD (
    echo ** build backend **
    go build -o server.exe backend\main.go
    echo !! build fininshed !!
)
goto :EOF

:usage
    echo build script for frontend and backend (default without args)
    echo use the following arguments (only one) to override default:
    echo -----------------------------------------------------------
    echo web - build the frontend code only
    echo run - build and run the backend code only
    echo runall - build the frontend and backend code and run the backend

:EOF