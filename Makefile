BINARY_NAME=crud

build:
	@go mod tidy
	@echo "Building CyberCRUD Binary..."
	@go run  main.go
	@go build -o tmp/${BINARY_NAME} .
	@go run ./mg  main.go

    # @go build -o binary/${BINARY_NAME}
	@echo "CRUD Binary built!"


run: build
	@echo "Starting Celeritas..."
	@./tmp/${BINARY_NAME} &
	@echo "Celeritas started!"
