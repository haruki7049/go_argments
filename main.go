package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) < 1{
		fmt.Print("no argments")
		os.Exit(1)
	}
	
	switch os.Args[1]{
	case "foo":
		fmt.Print("foo!!")
	case "unko":
		fmt.Print("unko.")
	case "hogehoge":
		fmt.Print("hogehoge~")
	default:
		fmt.Print("unknown argment")
	}
}
