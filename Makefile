BINARY_NAME=stickers
download:
	go mod download
build:
	go build -o ${BINARY_NAME} app/*.go
run: build
	./${BINARY_NAME}