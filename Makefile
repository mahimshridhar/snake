BINARY_NAME=go-cli-snake
 
build:
	go build -o ${BINARY_NAME} main.go snake.go
 
run:
	go build -o ${BINARY_NAME} main.go snake.go
	./${BINARY_NAME}
 
clean:
	go clean
	rm ${BINARY_NAME}

dep:
	go mod download