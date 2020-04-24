package mlib

import "fmt"

type Player interface {
	Play(source string)
}

func Play(source, musicGenre string) {
	var p Player
	switch musicGenre {
	case "MP3":
		p = &MP3Player{}
	case "WAV":
		p = &WAVPlayer{}
	default:
		_, _ = fmt.Println("Unsupported music type", musicGenre)
		return
	}
	p.Play(source)
}
