package mlib

import (
	"fmt"
	"time"
)

type MP3Player struct {
	stat     int
	progress int
}

func (p *MP3Player) Play(source string) {
	_, _ = fmt.Println("Playing MP3 music", source)
	p.progress = 0
	for p.progress < 100 {
		// 假装在播放
		time.Sleep(100 * time.Millisecond)
		_, _ = fmt.Print(".")
		p.progress += 10
	}
	_, _ = fmt.Println("\nFinished playing", source)
}
