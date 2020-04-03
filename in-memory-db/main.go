package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func e(err error) {
	log.Fatalln(err)
}

func handle(conn net.Conn) {
	defer conn.Close()

	io.WriteString(conn, "\r\nIN-MEMORY DATABASE\r\n\r\n"+
		"USE:\r\n"+
		"\tSET key value \r\n"+
		"\tGET key \r\n"+
		"\tDEL key \r\n\r\n"+
		"EXAMPLE:\r\n"+
		"\tSET fav chocolate \r\n"+
		"\tGET fav \r\n\r\n\r\n")

	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		if len(fs) < 1 {
			continue
		}
		switch fs[0] {
		case "GET":
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn, "%s\r\n", v)
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(conn, "EXPECTED VALUE\r\n")
				continue
			}
			k := fs[1]
			v := fs[2]
			data[k] = v
			fmt.Fprintf(conn, "%s = %s\r\n", k, v)
		case "DEL":
			k := fs[1]
			delete(data, k)
			fmt.Fprintf(conn, "DELETE %s\r\n", k)
		default:
			fmt.Fprintln(conn, "INVALID COMMAND "+fs[0]+"\r\n")
			continue
		}

	}
}

func main() {
	li, err := net.Listen("tcp", ":8080")
	e(err)
	defer li.Close()

	// listen for connections
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}

}