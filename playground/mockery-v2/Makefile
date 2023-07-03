all: generate test

dep:
	@echo 'Install dependencies'
	go get github.com/vektra/mockery/v2

generate:
	@echo 'Generate mocks'
	go generate ./...

test:
	go test ./...
