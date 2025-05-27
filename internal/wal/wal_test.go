package wal

import (
	"testing"

	protoapi "github.com/shohagrana006/simplewal/proto"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestAdd(t *testing.T) {
    result := Add(3, 3)
    expected := 6
    require.Equal(t,expected, result)
}

func TestWal(t *testing.T){
    result,err := Wal()

    entry := &protoapi.Entry{
		Name: "shohag",
	}
	expected, err1 :=  proto.Marshal(entry);

    require.NoError(t,err1)
    require.NoError(t,err)
    require.Equal(t,expected, result)
}