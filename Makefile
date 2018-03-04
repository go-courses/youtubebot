TARGET=telegram_bot

all: clean test build

clean:
	rm -rf $(TARGET)

test:
	go test ./...

build:
	go build -o $(TARGET) main.go
