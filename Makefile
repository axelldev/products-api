build:
	@go build -o bin/products

run: build
	@./bin/products

tests:
	@go test -v ./...