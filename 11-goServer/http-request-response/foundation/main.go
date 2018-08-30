package main

import (
	"bufio"
	"fmt"
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
	//read request
	request(conn)
	//write response
	response(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			//request line
			m := strings.Fields(ln)[0]
			fmt.Println("##Method -->", m)
		}
		if ln == "" {
			//header are done for
			break
		}
		i++
	}

}

func response(conn net.Conn) {
	body := "<html>" +
		"<head>" +
		"<meta charset=`UTF-8`" +
		"<title>go-test</title>" +
		"</head>" +
		"<body>" +
		"<h1>Hello</h1>" +
		"</body>" +
		"</html>"
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprint(conn, "Content-Length: \r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
