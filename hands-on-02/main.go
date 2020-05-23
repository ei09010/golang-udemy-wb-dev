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
		log.Fatal(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()

		if err != nil {
			log.Panic(err)
		}

		go handle(conn)
	}
}


func handle(conn net.Conn){
	defer conn.Close()
	i := 0
	var xString string
	var mString string
	scanner := bufio.NewScanner(conn)
		for scanner.Scan(){
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0{
			xString = strings.Fields(ln)[1]
			fmt.Println("our request string is: ", xString)
		}

		if i == 1{
			mString = strings.Fields(ln)[1]
			fmt.Println("our request string is: ", mString)

			fmt.Println("our full request is: ",mString,xString)
		}

		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}