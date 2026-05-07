package models

import "fmt"

type Byte uint8

func NewByte(v int) Byte {
	if v < 0 || v > 255 {
		panic(fmt.Sprintf("value %d out of range 0..255", v))
	}
	return Byte(v)
}