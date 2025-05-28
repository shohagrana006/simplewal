package wal

import (
	"encoding/binary"
	"io"
	"os"
	"path/filepath"

	pb "github.com/shohagrana006/simplewal/proto"
	"google.golang.org/protobuf/proto"
)

const (
	LogDir = "./log"
)

type WAL struct {
	file *os.File
}

func NewWAL() (*WAL, error){
	if err := os.MkdirAll(LogDir, 0755); err != nil{
		return nil, err
	}

	filename := filepath.Join(LogDir, "wal.log")

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil{
		return nil, err
	}

	w := &WAL{file: file}
	return w, nil
}

func (w *WAL) Write(entry *pb.Entry) error{
	data, err := proto.Marshal(entry)
	if nil !=  err{
		return err
	}

	lengthBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBuf, uint32(len(data)))

	if _, err := w.file.Write(lengthBuf); err != nil {
		return err
	}

	if _, err := w.file.Write(data); err != nil{
		return err
	}
	return nil
}

func ReadAll() ([]*pb.Entry, error){
	var entries []*pb.Entry
	filename := filepath.Join(LogDir, "wal.log")

	file, err := os.Open(filename)
	if err != nil{
		return nil, err
	}
	defer file.Close()

		for{
			lengthBuf := make([]byte, 4)
			_, err := io.ReadFull(file, lengthBuf)
			if err == io.EOF{
				break
			}
			if err != nil{
				return nil, err
			}

			length := binary.BigEndian.Uint32(lengthBuf)

			data := make([]byte, length)
			_, err = io.ReadFull(file, data)
			if err != nil{
				return nil, err
			}

			var entry pb.Entry
			if err := proto.Unmarshal(data, &entry); err != nil{
				return nil, err
			}
			entries = append(entries, &entry)
		}

	return entries, nil
}