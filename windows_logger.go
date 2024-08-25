//go:build windows

package main

import (
	"time"

	"golang.org/x/sys/windows"
)

const (
	keyCode = 0x13 // "Pause" key code (https://learn.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes)
)

var (
	user32           = windows.NewLazySystemDLL("user32.dll")
	getAsyncKeyState = user32.NewProc("GetAsyncKeyState")
)

func GetAsyncKeyState(vKey int) bool {
	ret, _, _ := getAsyncKeyState.Call(uintptr(vKey))
	return ret == 0x8001 || ret == 0x8000
}

func keyListener(ch chan int) {
	for {
		if GetAsyncKeyState(keyCode) {
			ch <- int(keyCode)
			time.Sleep(1000 * time.Millisecond)
		}
		time.Sleep(40 * time.Millisecond)
	}
}
