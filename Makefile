build:
	go build

clean:
	go clean

lint:
	golangci-lint run .

run:
	go run . $(OTHER)

setup:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.63.4

.PHONY: build clean lint run setup
