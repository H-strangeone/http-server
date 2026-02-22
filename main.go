package main

import (
	"bufio"
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
	reader:= bufio.NewReader(f)
	for {
		// data:= make([]byte,1024)
		// _,_ := f.Read(data)
		line, err:= reader.ReadBytes('\n')
		if err != nil {
			break
		}
		if err == io.EOF {
			break
		}
		fmt.Printf("%s", string(line))
		// fmt.Printf("%s", string(data[:n]))	
	}

}