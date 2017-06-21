package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"strings"
	"text/template"
)

const (
	methodGET  = "GET"
	methodPOST = "POST"
)

var (
	templateContent = `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>{{if .}}{{.Title}}{{end}}</title>
		</head>
		<body>
		{{if .}}
			<h1>"{{.Title}}"</h1>
		{{end}}
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
		{{if .}}
			{{.Content}}
		{{end}}
		</body>
		</html>
	`
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

type dataType struct {
	Title   string
	Content string
}

func handle(cnn net.Conn) {
	var method, URI string
	defer cnn.Close()
	scanner := bufio.NewScanner(cnn)

	i := true

	for scanner.Scan() {
		ln := scanner.Text()
		if i {
			fields := strings.Fields(ln)
			method = fields[0]
			URI = fields[1]
			i = !i
		}
		fmt.Println(ln)

		if ln == "" {
			break
		}
	}

	data := dataType{
		Title: method + " " + URI,
	}

	switch {
	case method == methodGET && URI == "/apply":
		data.Content = `
		<form action="/apply" method="POST">
			<input type="hidden" value="In my good death">
			<input type="submit" value="submit">
		</form>
		`
	case method == methodGET && URI == "/":
	case method == methodPOST && URI == "/apply":
	default:
	}

	tpl := template.New("template")
	tpl.Parse(templateContent)
	body := bytes.NewBufferString("")
	tpl.ExecuteTemplate(body, "template", data)
	result, _ := ioutil.ReadAll(body)
	fmt.Println(string(result))
	servContent(cnn, string(result))

	fmt.Println("Closing cnn.")
}

func servContent(cnn net.Conn, body string) {
	io.WriteString(cnn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(cnn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(cnn, "Content-Type: text/html\r\n")
	io.WriteString(cnn, "\r\n")
	io.WriteString(cnn, body)
}
