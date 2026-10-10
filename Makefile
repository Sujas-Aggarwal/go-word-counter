build:
	go build -o ./bin/ ./...

run:
	go run ./cmd/word-counter/main.go ${ARGS}

build-run: build
	./bin/word-counter ${ARGS}

seed: build
	./bin/dummy-data-creator ${ARGS}

test:
	go test -v ./...