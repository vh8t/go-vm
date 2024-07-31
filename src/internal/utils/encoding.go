package utils

import "encoding/binary"

func BytesToInt(bytes []byte) int64 {
	value := binary.LittleEndian.Uint64(bytes)
	return int64(value)
}

func IntToBytes(integer int64) []byte {
	value := make([]byte, 8)
	binary.LittleEndian.PutUint64(value, uint64(integer))
	return value
}
