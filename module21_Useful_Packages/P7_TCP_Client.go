package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func client(){
	client, err := net.Dial("tcp", "127.0.0.1:9090") // 127.0.0.0 is local host
	if err != nil{
		fmt.Println(err)
		return
	}
	message_to_send := "Hello, Fahad"
	fmt.Println("Sending message")
	err = gob.NewEncoder(client).Encode(message_to_send)
	if err != nil{
		fmt.Println(err)
	}
	client.Close()
}
func main(){
client()
}
