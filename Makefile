APP_NAME=thumbnail_previewer_backend

.PHONY: build run clean

default: run

build:
	@GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap api/main.go 
	@mkdir -p bin/
	@zip bin/thumbnail_previewer_backend.zip bootstrap
	@rm bootstrap

build-internal:
	@GOOS=linux GOARCH=amd64 go build -o main api/main.go

run:
	@go run api/main.go

clean:
	@rm -f $(APP_NAME)