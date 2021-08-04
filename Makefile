fmt:
	goimports -w ./
	go vet ./...
	go fmt ./...

run:
	go run ./cmd/main.go
