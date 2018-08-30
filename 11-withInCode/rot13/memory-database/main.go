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
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		go handle(conn)
	}

}

func handle(conn net.Conn) {
	defer conn.Close()
	io.WriteString(conn, "\n In- Memory Database\n"+
		" USE:\n SET key value \n GET key \n DEL key \n ")

	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		switch fs[0] {
		case "GET":
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn, v)

		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(conn, "Expected Value")
				continue
			}
			k := fs[1]
			v := fs[2]
			data[k] = v

		case "DEL":
			k := fs[1]
			delete(data, k)

		default:
			fmt.Fprintln(conn, "Invalid command")
		}
	}
}
