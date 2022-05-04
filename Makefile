BIN	=	ToDo

SRC	=	src/*.go

all:
	@go build -o $(BIN) $(SRC)
	@echo "Build binary !"
	@mv $(BIN) bin

run:
	@bin/./$(BIN) $(filter-out $@,$(MAKECMDGOALS))
%:
	@true

clean:
	@go clean
	@rm bin/$(BIN)
	@echo "Clean binary !"

re: clean all
