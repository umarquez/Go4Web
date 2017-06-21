package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
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

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if ln == "" {
			break
		}
	}

	body := "<!DOCTYPE html>"
	body += "<html>"
	body += "<body>"
	body += "<strong>HOLY COW</strong> THIS IS LOW LEVEL"
	body += "</body>"
	body += "</html>"
	io.WriteString(cnn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(cnn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(cnn, "Content-Type: text/html\r\n")
	io.WriteString(cnn, "\r\n")
	io.WriteString(cnn, body)

	fmt.Println("Closing cnn.")
}
