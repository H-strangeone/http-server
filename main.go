package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main(){
	f,err :=os.Open("messages.txt")
	if err!= nil{
		log.Fatal("Error creating file:", err)
	}
	defer f.Close() 
	for {
		data:= make([]byte, 8)
		n, err := f.Read(data)
		if err != nil {
			break
		}
		if err == io.EOF {
			break
		}
		if n>0{
			fmt.Printf("Read %d bytes: %s\n", n, string(data[:n]))	
		}
	}

}