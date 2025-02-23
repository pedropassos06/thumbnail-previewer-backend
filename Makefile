APP_NAME=thumbnail_previewer_backend

.PHONY: build run clean

default: run

build:
	@GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap api/main.go 
	@mkdir -p bin/
	@zip bin/thumbnail_previewer_backend.zip bootstrap
	@rm bootstrap

run:
	@go run api/main.go

clean:
	@rm -f $(APP_NAME)