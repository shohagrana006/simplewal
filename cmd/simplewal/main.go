package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/shohagrana006/simplewal/internal/wal"
	pb "github.com/shohagrana006/simplewal/proto"
)

func main(){
	walWriter, err := wal.NewWAL()
	if err != nil{
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter key: ")
	key, _ := reader.ReadString('\n')
	key = strings.TrimSpace(key)

	fmt.Println("Enter your value: ")
	value, _ := reader.ReadString('\n')
	value = strings.TrimSpace(value)

	entry := &pb.Entry{
		Key: key,
		Value: value,
	}
	
	if err := walWriter.Write(entry); err != nil{
		panic(err)
	}
	

	entries, err := wal.ReadAll()
	if err != nil{
		panic(err)
	}

	for _, e := range entries{
		fmt.Printf("Key is: %s & Value is: %s\n", e.Key, e.Value)
	}



}