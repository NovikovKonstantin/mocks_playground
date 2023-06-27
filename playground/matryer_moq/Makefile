all: generate test

dep:
	@echo 'Install dependencies'
	go install github.com/matryer/moq@latest

generate:
	@echo 'Generate mocks'
	go generate ./...

test:
	go test ./...
