migration:
	go run cmd/migration/main.go

seed:
	go run cmd/seed/main.go

run.app:
	go build -o tmp/main cmd/app/main.go
	./tmp/main

example:
	gow run examples/app/main.go
