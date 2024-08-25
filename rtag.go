package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int)
	os := runtime.GOOS

	switch os {
	case "linux":
		keyboards := findKeyboardsFromSysClass()
		for _, keyboard := range keyboards {
			go linuxKeyListener(keyboard, ch)
		}
	case "windows":
		go windowsKeyListener(ch)
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
