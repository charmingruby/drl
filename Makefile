build:
	go build -o ./bin/example ./example/main.go

run: build
	./bin/example
