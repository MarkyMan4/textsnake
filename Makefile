BIN=textsnake

build:
	go build -o $(BIN) .

clean:
	rm $(BIN)
