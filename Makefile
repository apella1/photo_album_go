.DEFAULT_GOAL := run
.PHONY = format vet build run

format:
	@echo "Formatting code..."
	@go fmt ./...

vet: format
	@echo "Vetting code..."
	@go vet ./...

build: vet
	@echo "Building..."
	@go build -o main cmd/api/main.go
run: build
	./main