all: generate test

fmt:
	@echo 'Run goimports'
	find . -name "*.go" -exec goimports -w {} \;

dep:
	@echo 'Install dependencies'
	go install github.com/gojuno/minimock/v3/cmd/minimock@v3.1.3

generate:
	@echo 'Generate mocks'
	go generate ./...

test:
	go test ./...
