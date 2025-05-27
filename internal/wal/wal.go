package wal

import (
	"google.golang.org/protobuf/proto"
	protoapi "github.com/shohagrana006/simplewal/proto"
)



func Add(a,b int) int{
	return a+b
}

func Wal() ([]byte, error) {

	entry := &protoapi.Entry{
		Name: "shohag",
	}
	return proto.Marshal(entry);
}
