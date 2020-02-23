package utils

import (
	"bytes"
	"encoding/binary"
)

// IntToBytes int类型转换成bytes，采用大端模式
func IntToBytes(i int) []byte {
	buffer := bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.BigEndian, i)
	return buffer.Bytes()
}

// BytesToInt bytes转化成int，采用大端模式
func BytesToInt(b []byte) int {
	var i int
	buffer := bytes.NewBuffer(b)
	binary.Read(buffer, binary.BigEndian, &i)
	return i
}

// IntToBytes int64类型转换成bytes，采用大端模式
func Int64ToBytes(i int64) []byte {
	buffer := bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.BigEndian, i)
	return buffer.Bytes()
}

// BytesToInt bytes转化成int64，采用大端模式
func BytesToInt64(b []byte) int64 {
	var i int64
	buffer := bytes.NewBuffer(b)
	binary.Read(buffer, binary.BigEndian, &i)
	return i
}