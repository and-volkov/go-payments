build:
	go build -o bin/gopayments

run: build
	./bin/gopayments

test:
	go test -v ./...
