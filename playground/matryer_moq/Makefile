all: generate test

fmt:
	@echo 'Run goimports'
	find . -name "*.go" -exec goimports -w {} \;

dep:
	@echo 'Install dependencies'
	go install github.com/matryer/moq@latest

generate:
	@echo 'Generate mocks'
	go generate ./...

test:
	go test ./...
