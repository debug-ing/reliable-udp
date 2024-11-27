APP_NAME := reliable-udp
SRC_SERVER := server/main.go
SRC_CLIENT := client/main.go

server:
	@echo "Running server..."
	go run $(SRC_SERVER)

client:
	@echo "Running client..."
	go run $(SRC_CLIENT)

help:
	@echo "Makefile targets:"
	@echo "  server     - Run server file"
	@echo "  client     - Run client file"