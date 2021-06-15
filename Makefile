lint:
	golangci-lint run

fmt:
	goimports -w ./
	go vet ./...
	go fmt ./...
	make lint

run:
	go run ./playground/cmd/main.go
