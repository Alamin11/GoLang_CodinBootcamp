package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//Listening to the request
	nl, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	//Accepting the request
	for {
		conn, err := nl.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		bs := make([]byte, 1024) //a byte slice to store the request
		//reading the request as byte string
		n, err := conn.Read(bs) //kotha shona
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(n)

		//converting byte string bs to normal string
		reqStr := string(bs) //conversion (between two diffrent types)
		fmt.Println(reqStr)

		body := `<!DOCTYPE html><html lang ="en"><head><title>networking</title></head><body><h1>This is a testing string <br> <p>New string added</p></h1></body></html>`

		resp := "HTTP/1.1 200 OK\r\n" +
			"Content-Length: %d\r\n" +
			"Content-Type: text/html\r\n" +
			"\r\n%s"

		msg := fmt.Sprintf(resp, len(body), body)
		fmt.Println(msg)        //displaying response message
		conn.Write([]byte(msg)) //kotha bola

	}
}
