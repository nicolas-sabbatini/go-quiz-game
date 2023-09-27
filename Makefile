BINARYNAME=quizes
PATH_TO_BINARY=./cmd/quizes
.DEFAULT_GOAL := run

run:
	go run ${PATH_TO_BINARY}/${BINARYNAME}.go

build-linux:
	GOARCH=amd64 GOOS=linux go build -o target/${BINARYNAME}-linux-amd64 ${PATH_TO_BINARY}/${BINARYNAME}.go

build-windows:
	GOARCH=amd64 GOOS=windows go build -o target/${BINARYNAME}-windows-amd64.exe ${PATH_TO_BINARY}/${BINARYNAME}.go

build-mac:
	GOARCH=amd64 GOOS=darwin go build -o target/${BINARYNAME}-darwin-amd64 ${PATH_TO_BINARY}/${BINARYNAME}.go

build-all: build-linux build-windows build-mac

build: build-linux

clear:
	rm -rf target