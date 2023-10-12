BIN=textsnake

build:
	go build -o $(BIN) .

buildall:
	env GOOS=linux GOARCH=amd64 go build -o $(BIN)_linux_amd64 .
	env GOOS=linux GOARCH=arm64 go build -o $(BIN)_linux_arm64 .
	env GOOS=darwin GOARCH=amd64 go build -o $(BIN)_darwin_amd64 .
	env GOOS=darwin GOARCH=arm64 go build -o $(BIN)_darwin_arm64 .
	env GOOS=windows GOARCH=amd64 go build -o $(BIN)_windows_amd64.exe .
	env GOOS=windows GOARCH=arm64 go build -o $(BIN)_windows_arm64.exe .

clean:
	rm $(BIN)_*
