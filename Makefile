
get-deps:
	go get .
build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -ldflags "-w -s -X github.com/tibbyrocks/tibby/internal/commands/tibbycmds.version=$(shell git describe --tags --always --long)" -o bin/tibby_linux_amd64
build-linux-arm64:
	GOOS=linux GOARCH=arm64 go build -ldflags "-w -s -X github.com/tibbyrocks/tibby/internal/commands/tibbycmds.version=$(shell git describe --tags --always --long)" -o bin/tibby_linux_arm64
build-windows-amd64:
	GOOS=windows GOARCH=amd64 go build -ldflags "-w -s -X github.com/tibbyrocks/tibby/internal/commands/tibbycmds.version=$(shell git describe --tags --always --long)" -o bin/tibby_windows_amd64.exe

build-all: build-linux-amd64 build-linux-arm64 build-windows-amd64