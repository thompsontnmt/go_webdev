package main

import (
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

for {
	c, err := li.Accept()
	if err != nil {
		log.Println(err)
		}
		io.WriteString(c, "I see you connected")

		c.Close()
	}
}

