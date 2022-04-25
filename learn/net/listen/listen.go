package main

import (
	"bytes"
	"log"
	"net"
	"time"
)

func main() {
	go listen()
	time.Sleep(2 * time.Second)
	go write()
	select {}
}
func write() {
	buf := make([]byte, 1)
	buf[0] = '1'
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Println("err: ", err.Error())
		return
	}
	defer conn.Close()

	conn.Write([]byte("string"))
	log.Println("write-success!")
}
func listen() {
	var err error
	buf := make([]byte, 1024)
	buffer := bytes.Buffer{}
	lis, _ := net.Listen("tcp", "127.0.0.1:8080")
	conn, _ := lis.Accept()
	n, err := conn.Read(buf)
	if err != nil {
		log.Println(err.Error())
		return
	}
	buffer.Write(buf[:n])
	log.Println("reading: " ,buffer.String())
}
