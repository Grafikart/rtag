package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go keyListener(ch)
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
