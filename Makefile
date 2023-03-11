GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

.iPHONE: chatgpt-example
chatgpt-example:
	cd example
	go build -o chatgpt
