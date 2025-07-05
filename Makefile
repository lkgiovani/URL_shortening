build:
	go build -o bin/main cmd/main.go

run: build
	./bin/main


gomod:
	go mod tidy
