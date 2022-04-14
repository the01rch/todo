include config.mk

build:
	@$(CC) -o $(BIN) $(FILE)
	@echo "Build binary !"
	@chmod 755 $(BIN)

run:
	@go run $(FILE)

clean:
	@rm $(BIN)
	@echo "Clean binary !"

re: clean build
