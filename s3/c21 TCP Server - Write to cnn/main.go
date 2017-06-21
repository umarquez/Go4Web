package main

import (
	"fmt"
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
		}

		fmt.Fprintln(cnn, "Hello world!")

		err = cnn.Close()
		if err != nil {
			log.Println(err)
		}
	}
}
