BIN	=	todo

SRC	=	src/*.go

all:
	gccgo -o $(BIN) $(SRC)
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
