package main

import (
	"syscall"
	"unsafe"
)

// eventSize is size of structure of InputEvent (24 bytes)
var eventSize = int(unsafe.Sizeof(InputEvent{}))

// InputEvent is the keyboard event structure itself
type InputEvent struct {
	Time  syscall.Timeval
	Type  uint16
	Code  uint16
	Value int32
}

func (i *InputEvent) KeyPress() bool {
	return i.Value == 1
}

func (i *InputEvent) KeyRelease() bool {
	return i.Value == 0
}
