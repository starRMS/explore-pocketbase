package utils

import "unsafe"

func ConvertStringToSliceByte(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
