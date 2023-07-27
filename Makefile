hello:
	echo "Hello"

build:
	go build -o bin/main main.go

run:
	go run main.go

test:
	go test ./...

coverage:
	go test -coverpkg=./... ./...
