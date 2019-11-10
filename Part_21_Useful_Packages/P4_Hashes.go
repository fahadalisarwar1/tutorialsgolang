package main

import (
	"fmt"
	"hash/crc32"
)

/// in programming a hash is a function that takes a data and reduces it to a smaller fiixed size.

// Two type of hash functions in GO cryptographic and non cryptographic

func main(){
	// create a hash function

	hash :=  crc32.NewIEEE()
	hash.Write([]byte("My name is fahad"))


	// calculate checksum
	checksum := hash.Sum32()
	fmt.Println(checksum)
	// checksum is used to verify the integrity of files, if two files have same checksum, it is highly probable that these two files are same and havent
	// been modified

}