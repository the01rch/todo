BIN	=	todo

SRC	=	src/*.go

all:
	#gccgo -o $(BIN) $(SRC)
	@go build -o $(BIN) $(SRC)
	@echo "Build binary !"
	@sudo cp todo /bin/

clean:
	@go clean
	@rm $(BIN)
	@sudo rm /bin/todo
	@echo "Clean binary !"

install: all
	@sudo cp data/todo.json ~/.todo.json
	@sudo chmod a+rw ~/.todo.json
	@echo "Installation done !"

re: clean all
