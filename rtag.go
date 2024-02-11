package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	keyboards := findKeyboardsFromSysClass()
	ch := make(chan int)
	for _, keyboard := range keyboards {
		go keyListener(keyboard, ch)
	}
	fmt.Println("Listening for \"Pause\" key...")
	var start time.Time
	var subtitle *Subtitle
	for range ch {
		// La première pression démarre le timer
		if start.IsZero() {
			start = time.Now()
			fmt.Println("Timer is starting")
		}
		if subtitle == nil {
			subtitle = MakeSubtitle("cuts.srt", start)
		}
		subtitle.AddMarker()
		fmt.Printf("Marker added at %s\n", timeCode(time.Now().Sub(start)))
	}
}

func keyListener(path string, ch chan int) {
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeCharDevice)
	if err != nil {
		panic(err)
	}
	for {
		buffer := make([]byte, eventSize)
		n, err := f.Read(buffer)
		if err != nil {
			panic(err)
		}
		if n <= 0 {
			return
		}
		event, err := eventFromBuffer(buffer)
		// Listen for the "Pause" key
		if err == nil && event.KeyPress() && event.Code == 119 {
			ch <- int(event.Code)
		}
	}

}

func eventFromBuffer(buffer []byte) (*InputEvent, error) {
	event := &InputEvent{}
	err := binary.Read(bytes.NewBuffer(buffer), binary.LittleEndian, event)
	return event, err
}

func findKeyboardsFromInput() (inputs []string) {
	files, err := os.ReadDir("/dev/input")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			panic(err)
		}
		if info.Mode()&os.ModeCharDevice != 0 {
			inputs = append(inputs, file.Name())
		}
	}
	return inputs
}

func findKeyboardsFromSysClass() []string {
	var devices []string
	paths, err := filepath.Glob("/sys/class/input/event*")
	if err != nil {
		panic(err)
	}

	for _, path := range paths {
		content, err := os.ReadFile(path + "/device/name")
		if err != nil {
			panic(err)
		}
		deviceName := strings.ToLower(string(content))
		if strings.Contains(deviceName, "keyboard") {
			devices = append(devices, strings.Replace(path, "/sys/class/input", "/dev/input", 1))
		}
	}
	return devices
}
