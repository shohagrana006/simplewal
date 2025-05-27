package main

import (
	"fmt"

	"github.com/shohagrana006/simplewal/internal/wal"
)

func main(){
	result := wal.Add(2,3)
	fmt.Printf("hello test build file %v", result)
}