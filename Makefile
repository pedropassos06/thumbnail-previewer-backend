APP_NAME=thumbnail-previewer-backend

.PHONY: build run clean

build:
	@go build -o $(APP_NAME)

run:
	@go run main.go

clean:
	@rm -f $(APP_NAME)