package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main(){
	func_to_pass := func(path string, info os.FileInfo, err error)(error){
		fmt.Println(path)
		return nil
	}
	// This function is called for every file and folder in the directory
	//  path , which is the path to the file; info ,
	//which is the information for the file (the same information you get from using
	//os.Stat ); and err , which is any error that was received while walking the directory.
	//The function returns an error and you can return filepath.SkipDir to stop walking
	//immediately.
	filepath.Walk(".", func_to_pass)
}
