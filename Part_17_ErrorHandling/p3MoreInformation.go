package main

import (
	"fmt"
	"net"
)

// The second way to get more information is to assert for the underlying type and get more
// information by calling methods on the struct type.

// Dns Errors

func main(){
	addr, err := net.LookupHost("googleeeeeeeeeeeeeeeeee.com")
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println("temporary error")
		} else {
			fmt.Println("generic error: ", err)
		}
		return
	}
	fmt.Println(addr)
}

