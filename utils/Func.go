package utils

import "unsafe"

func Byte2Str(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}
