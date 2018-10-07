.PHONY: run
run: format
	go run example/main.go ./example/hoge.formula

.PHONY: build
build: format
	go build

.PHONY: format
format:
	goimports -w .
	gofmt -w .

