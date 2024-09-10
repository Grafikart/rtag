package main

import (
	"fmt"
	"os"
	"time"
)

// Method to write markers to a .srt file
type Subtitle struct {
	start time.Time
	index int32
	file  *os.File
}

func MakeSubtitle(path string, start time.Time) *Subtitle {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	// On linux we need to start the app using sudo, ensure the file is owned by the user
	os.Chown(path, 1000, 1001)
	return &Subtitle{
		start: start,
		file:  f,
	}
}

func (s *Subtitle) AddMarker() {
	elapsedTime := time.Since(s.start)
	s.index++
	s.file.WriteString(fmt.Sprintf("%d\n%s --> %s\nCUT\n\n", s.index, timeCode(elapsedTime), timeCode(elapsedTime+time.Second)))
}

func (s *Subtitle) Close() {
	s.file.Close()
}

func timeCode(d time.Duration) string {
	hours := d / time.Hour
	d -= hours * time.Hour
	minutes := d / time.Minute
	d -= minutes * time.Minute
	seconds := d / time.Second
	return fmt.Sprintf("%02d:%02d:%02d,000", hours, minutes, seconds)
}
