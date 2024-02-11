package main

import (
	"fmt"
	"os"
	"time"
)

// Method to write markers to a .srt file
type Subtitle struct {
	start time.Time
	file  *os.File
}

func MakeSubtitle(path string, start time.Time) *Subtitle {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	err = os.Chown(path, 1000, 1001)
	if err != nil {
		panic(err)
	}
	return &Subtitle{
		start: start,
		file:  f,
	}
}

func (s *Subtitle) AddMarker() {
	elapsedTime := time.Now().Sub(s.start)
	s.file.WriteString(fmt.Sprintf("%s --> %s\nCUT\n\n", timeCode(elapsedTime), timeCode(elapsedTime+time.Second)))
	return
}

func (s *Subtitle) Close() {
	s.file.Close()
	return
}

func timeCode(d time.Duration) string {
	hours := d / time.Hour
	d -= hours * time.Hour
	minutes := d / time.Minute
	d -= minutes * time.Minute
	seconds := d / time.Second
	return fmt.Sprintf("%02d:%02d:%02d,000", hours, minutes, seconds)
}
