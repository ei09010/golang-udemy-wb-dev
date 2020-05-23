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

	lt, err := net.Listen("tcp", ":8080")

	if err != nil{
		log.Fatal(err)
	}

	defer lt.Close()


	for{
		conn, err := lt.Accept()

		if err != nil{
			log.Println(err)
			continue
		}

			go serve(conn)

		// we never get here
		// we have an open stream connection
		// how does the above reader know when it's done?
		fmt.Println("Code got here.")

	}

}

func response(c net.Conn, method string, uriString string) {

	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Code Gangsta</title>
		</head>
		<body>
			<h1>"HOLY COW THIS IS LOW LEVEL"</h1>
		</body>
		</html>
	`

	body += "\n"
	body += method
	body += "\n"
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}



func serve(conn net.Conn){

	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	var methodString string
	var uriString string
	var i int

	for scanner.Scan(){
		ln := scanner.Text()
		scanParts := strings.Fields(ln)

		if i == 0{
			methodString = scanParts[0]
			uriString = scanParts[1]


			fmt.Println(methodString,"\n",uriString)

			methodUriString := methodString+uriString
			if methodUriString == "GET /" {
				response(conn, methodString, uriString)
			}

		}
		fmt.Println(ln)

		if ln == ""{
			break
		}

		i++
	}
}
