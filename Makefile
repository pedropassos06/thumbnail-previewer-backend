APP_NAME=thumbnail_previewer_backend

.PHONY: build run clean

default: build

build:
	@GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap api/main.go 
	@mkdir -p bin/
	@zip bin/$(APP_NAME).zip bootstrap
	@mv bootstrap bin/$(APP_NAME)

run:
	@sam local start-api --env-vars env.json

clean:
	@rm -rf bin/

test:
	@go test -v ./...