run:
	gofumpt -l -w .
	go run cmd/main.go
build:
	gofumpt -l -w .
	go build -o cmd/main.go 
