package mlib

import (
	"fmt"
	"time"
)

type WAVPlayer struct {
	stat     int
	progress int
}

func (p *WAVPlayer) Play(source string) {
	_, _ = fmt.Println("Playing WAV music", source)
	p.progress = 0
	for p.progress < 100 {
		// 假装在播放
		time.Sleep(100 * time.Millisecond)
		_, _ = fmt.Print(".")
		p.progress += 10
	}
	_, _ = fmt.Println("\nFinished playing", source)
}
