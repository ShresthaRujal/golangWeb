package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\nHello from TCP server \n")

		//Fprintln takes the writer and write back
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintln(conn, "Well, I hope!")
		conn.Close()
	}
}
