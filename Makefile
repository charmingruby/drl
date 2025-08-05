BIN := drl

build:
	go build -o ./bin/$(BIN) ./cmd/drl/main.go

run: build
	./bin/$(BIN)
