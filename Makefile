include config.mk

build:
	@$(CC) -o $(BIN) $(FILE)
	@echo "Build binary !"

run:
	@go run $(FILE)

clean:
	@trash $(BIN)
	@echo "Clean binary !"

re: clean build
