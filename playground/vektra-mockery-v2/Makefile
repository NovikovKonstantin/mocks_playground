all: generate test

fmt:
	@echo 'Run goimports'
	find . -name "*.go" -exec goimports -w {} \;

dep:
	@echo 'Install dependencies'
	go get github.com/vektra/mockery/v2@v2.31.1

generate:
	@echo 'Generate mocks'
	go generate ./...

test:
	go test ./...
