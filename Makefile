build:
	go build -o ./bin/ ./...

run:
	go run ./cmd/word-counter/main.go

build-run: build
	./bin/word-counter

seed: build
	./bin/dummy-data-creator