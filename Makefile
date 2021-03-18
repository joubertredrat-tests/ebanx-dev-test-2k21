build-mac:
	env GOOS=darwin GOARCH=amd64 go build -o ebanx-darwin-amd64

build-linux:
	go build -o ebanx-linux-amd64
