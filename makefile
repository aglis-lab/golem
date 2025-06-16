.PHONY: example
example:
	gow run example/main.go

.PHONY: cli
cli:
	go build -o temp/golem cmd/cli/main.go
