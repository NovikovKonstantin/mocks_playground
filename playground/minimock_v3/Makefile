# Dummy makefile

all: generate test

dep:
	go get github.com/gojuno/minimock/v3/cmd/minimock
	go install github.com/gojuno/minimock/v3/cmd/minimock
	go mod tidy

generate:
	cd ./service && minimock
	go mod tidy

test:
	go test -v ./...
