package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		cnn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handle(cnn)
	}
}

func handle(cnn net.Conn) {
	defer cnn.Close()
	scanner := bufio.NewScanner(cnn)

	body := ""
	i := true
	for scanner.Scan() {
		ln := scanner.Text()
		if i {
			fields := strings.Fields(ln)
			body = "Method: " + fields[0] + "\r\n"
			body += "URI: " + fields[1] + "\r\n"
		}
		fmt.Println(ln)

		if ln == "" {
			break
		}
		i = false
	}

	fmt.Println("Code got here.")

	io.WriteString(cnn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(cnn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(cnn, "Content-Type: text/plain\r\n")
	io.WriteString(cnn, "\r\n")
	io.WriteString(cnn, body)

	fmt.Println("Closing cnn.")
}
