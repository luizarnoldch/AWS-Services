@echo off

setlocal

REM Eliminar los archivos main y main.zip si existen
if exist main (
  del main
)

if exist main.zip (
  del main.zip
)

REM Configurar las variables de entorno
set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0

REM Compilar el archivo main.go
go build main.go

REM Esperar a que se compile el código antes de continuar
if not errorlevel 1 (
  REM Crear el archivo main.zip utilizando build-lambda-zip.exe
  build-lambda-zip.exe -o main.zip main
)

endlocal
