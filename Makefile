init:
	rm -rf go.mod go.sum
	go mod init gonote
	go mod tidy
bin:init
	go build main.go

clean:
	rm -rf *.o main go.mod go.sum