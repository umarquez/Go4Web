package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

/*
 * Building upon the code from the previous exercise:
 * - In that previous exercise, we WROTE to the connection.
 * - Now I want you to READ from the connection.
 * - You can READ and WRITE to a net.Conn as a connection implements both the
 *   reader and writer interface.
 * - Use bufio.NewScanner() to read from the connection.
 * - After all of the reading, include these lines of code:
 * - fmt.Println("Code got here.") io.WriteString(c, "I see you connected.")
 * - Launch your TCP server.
 * - In your web browser, visit localhost:8080.
 * - Now go back and look at your terminal.
 * - Can you answer the question as to why "I see you connected." is never written?
 *
 * R: Because is an open stream, we never know where to stop.
 */

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
		}

		go handle(cnn)
	}
}

func handle(cnn net.Conn) {
	defer cnn.Close()
	scanner := bufio.NewScanner(cnn)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	fmt.Println("Code got here.")
	_, err := io.WriteString(cnn, "I see you connected.")
	if err != nil {
		log.Println(err)
	}
}
