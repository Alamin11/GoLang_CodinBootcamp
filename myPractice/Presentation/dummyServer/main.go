package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	nl, err := net.Listen("tcp", ":8089")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for {
		conn, err := nl.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		bs := make([]byte, 1024)

		n, err := conn.Read(bs)
		fmt.Println(n)

		req := string(bs)
		fmt.Println(req)

		body :=
			`
			<!DOCTYPE html>
			<html>
			<head><title>Dummy server</title></head>
			<body><h1>Welcome to my website</h1></body>
			</html>
			`
		resp := "HTTP/1.1 200 OK \r\n" +
			"Content-Length : %d  \r\n" +
			"Content-Type: text/html\r\n" +
			"\r\n %s"

		msg := fmt.Sprintf(resp, len(body), body)
		fmt.Println(msg)

		conn.Write([]byte(msg))
	}

}
