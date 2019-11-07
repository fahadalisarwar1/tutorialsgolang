package main
import (
"flag"
"fmt"
)

func main() {
	fptr := flag.String("fpath", "sample.txt", "file path to read from")
	flag.Parse()
	fmt.Println("value of fpath is", *fptr)
}
