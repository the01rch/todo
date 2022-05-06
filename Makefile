BIN	=	ToDo

SRC	=	src/*.go

all:
	@go build -o $(BIN) $(SRC)
	@echo "Build binary !"

run:
	@./$(BIN) $(filter-out $@,$(MAKECMDGOALS))
%:
	@true

clean:
	@go clean
	@rm $(BIN)
	@echo "Clean binary !"

re: clean all
