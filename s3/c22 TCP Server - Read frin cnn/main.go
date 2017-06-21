package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
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

		//fmt.Fprintln(cnn, "Waiting for an input.")

		go handle(cnn)
	}
}

func handle(cnn net.Conn) {
	defer cnn.Close()

	scnr := bufio.NewScanner(cnn)
	//scnr.Split(bufio.ScanWords)

	for scnr.Scan() {
		ln := scnr.Text()
		fmt.Println(ln)
	}

	fmt.Println("-- END --")
	// Kills the goroutine
	os.Exit(0)
}
