package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main(){


	// Regexp	Meaning
	// x*	zero or more x, prefer more
	// x*?	prefer fewer (non-greedy)
	// x+	one or more x, prefer more
	// x+?	prefer fewer (non-greedy)
	// x?	zero or one x, prefer one
	// x??	prefer zero
	// x{n}	exactly n x
	



	str := "my name is fahad"
	re := regexp.MustCompile(`f[a-z]{2}[a-z]+`) // good practice to use raw strings as search expressions
	res := re.FindAllString(str, -1)
	fmt.Println(res)

	ipconfig := "wlp6s0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500 " +
		"inet 192.168.0.12  netmask 255.255.255.0  broadcast 192.168.0.255" +
		"ether 18:1d:ea:ab:e2:48  txqueuelen 1000  (Ethernet)" +
		"RX packets 6474628  bytes 6767278434 (6.7 GB)" +
		"RX errors 0  dropped 0  overruns 0  frame 0" +
		"TX packets 4812283  bytes 5870101846 (5.8 GB)" +
		"TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0"

	// lets say we search the ip value from this string

	re = regexp.MustCompile(`inet [0-9]*.{1}[0-9]*.{1}[0-9]*.{1}[0-9]*`)
	results := re.FindAllString(ipconfig,-1)
	fmt.Println(strings.TrimPrefix(results[0], "inet "))





}
