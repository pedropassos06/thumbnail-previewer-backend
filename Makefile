APP_NAME=thumbnail_previewer_backend

.PHONY: build run clean

default: run

build:
	@go build -o $(APP_NAME)

run:
	@go run cmd/main.go

clean:
	@rm -f $(APP_NAME)