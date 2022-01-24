include .env

NAME=gin-examples
OUTPUT_DIR=output
TARGET=$(OUTPUT_DIR)/$(NAME)
LD_FLAGS=-s -w

.PHONY: build clean download run run-prod

build:
	go build -ldflags "$(LD_FLAGS)" -o $(TARGET) $(NAME)
	upx -9 -o $(TARGET)-upx-9 $(TARGET)
	@ls -lh --color $(OUTPUT_DIR)

clean:
	rm -rf $(OUTPUT_DIR)/*

download:
	go mod download

run:
	go run main.go

run-prod: build
	./$(TARGET)
