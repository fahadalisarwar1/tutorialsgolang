package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func getHashValue(filename string)(uint32, error){
	file, err := os.Open(filename)
	if err != nil{
		fmt.Println("Cant open file")
		return 0, err
	}
	defer file.Close()
	hash := crc32.NewIEEE()
	// copy the file into the hasher
	// - copy takes (dst, src) and returns (bytesWritten, error)
	_, err = io.Copy(hash, file)
	if err != nil{
		fmt.Println("Cant copy file to hasher")
		return 0, err
	}
	return hash.Sum32(), nil

}

func main(){
	h1, err :=  getHashValue("file_1")
	if err != nil{
		fmt.Println(err)
	}

	h2, err := getHashValue("file_2")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(h1, h2, h1==h2)
}
