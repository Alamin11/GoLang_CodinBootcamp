package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	nl, err := net.Listen("tcp", ":8088")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	conn, err := nl.Accept()
	if err != nil {
		fmt.Println(err.Error())
	}
	bs := make([]byte, 1024)
	n, err := conn.Read(bs)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(n)
	reqStr := string(bs)
	fmt.Println(reqStr)

	body := `
			<!DOCTYPE html>
			<html lang ="en">
			<head>
			<title>Dummy server</title>
			</head>
			<body>
				<h1>Wellcome to my website</h1>
				<p>Its a dummy webserver testing </p>
			</body>
			</html>
	`
	resp := "HTTP/1.1 200 OK \r\n" +
		"Content-Length : %d\r\n" +
		"Cntent-Type : text/html\r\n" +
		"\r\n %s"
	msg := fmt.Sprintf(resp, len(body), body)

	fmt.Println(msg)

	conn.Write([]byte(msg))

}
