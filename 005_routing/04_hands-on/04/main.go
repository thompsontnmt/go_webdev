package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
		}
		defer conn.Close()
	}
}