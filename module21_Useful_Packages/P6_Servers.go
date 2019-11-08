package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

// TCP SERVER

// TCP is the primary protocol used for communication over the Internet. Any time you
//interact with a web page, play a multiplayer computer game, stream a movie, or video
//chat, there’s a good chance your computer is communicating with a remote server
//using TCP.
//In Go, we can create a TCP server using the net package’s Listen function. Listen
//takes a network type (in our case, tcp ) and an address and port to bind, and returns a
//net.Listener :

func handleTCPConnection(conn net.Conn){
	var msg string
	err := gob.NewDecoder(conn).Decode(&msg)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Received Message: ", msg)
	}

}
func server(){
	ip := "localhost"
	port := "9090"
	address := ip + ":" + port
	listener,  err := net.Listen("tcp", address)
	if err != nil{
		fmt.Println(err)
		return
	}

	// create an infinite loop for a listener
	for {
		connection, err := listener.Accept()
		if err != nil{
			fmt.Println(err)
			continue
		}
		fmt.Println("Got a Connection from ", connection.RemoteAddr().String())
		defer connection.Close()
		go handleTCPConnection(connection)
	}
}
func main(){
	server()

}