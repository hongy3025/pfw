package main

 import (
 	"io"
 	"log"
	"net"
	"flag"
 )

var localAddr *string = flag.String("local", ":8080", "local address")
var remoteAddr *string = flag.String("remote", "127.0.0.1:11000", "remote address")

 func main() {

 	ln, err := net.Listen("tcp", *localAddr)
 	if err != nil {
 		log.Fatal(err)
 	}

 	log.Println("Port forwarding server up and listening on ", *localAddr)

 	for {
 		conn, err := ln.Accept()
 		if err != nil {
 			log.Fatal(err)
 		}

 		go handleConnection(conn)
 	}
 }

 func forward(src, dest net.Conn) {
 	defer src.Close()
 	defer dest.Close()
 	io.Copy(src, dest)
 }

 func handleConnection(c net.Conn) {
 	remote, err := net.Dial("tcp", *remoteAddr)
 	if err != nil {
		 log.Println(err)
		 return
 	}

 	go forward(c, remote)
 	go forward(remote, c)
 }