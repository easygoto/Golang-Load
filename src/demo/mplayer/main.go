package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"demo/mplayer/library"
	"demo/mplayer/mlib"
)

var lib *library.MusicManager
var id = 1

func handleLibCommands(tokens []string) {
	if len(tokens) < 2 {
		help()
		return
	}
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			_, _ = fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		if len(tokens) == 7 {
			id++
			lib.Add(&library.MusicEntry{Id: strconv.Itoa(id), Name: tokens[2], Artist: tokens[3], Source: tokens[4],
				Type: strings.ToUpper(tokens[5]), Genre: strings.ToUpper(tokens[6])})
		} else {
			_, _ = fmt.Println("USAGE: lib add <name> <artist> <source> <type> <genre>")
		}
	case "remove":
		if len(tokens) == 3 {
			_ = lib.RemoveByName(tokens[2])
		} else {
			_, _ = fmt.Println("USAGE: lib remove <name>")
		}
	default:
		_, _ = fmt.Println("Unrecognized lib command:", tokens[1])
	}
}

func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		_, _ = fmt.Println("USAGE: play <name>")
		return
	}
	e := lib.Find(tokens[1])
	if e == nil {
		_, _ = fmt.Println("The music", tokens[1], "does not exist.")
		return
	}
	mlib.Play(e.Source, e.Genre)
}

func help() {
	_, _ = fmt.Println(` 
 Enter following commands to control the player:
 lib list -- View the existing music lib
 lib add <name> <artist> <source> <type> <genre> -- Add a music to the music lib
 lib remove <name> -- Remove the specified music from the lib
 play <name> -- Play the specified music
 `)
}

func main() {
	help()
	lib = library.NewMusicManager()
	r := bufio.NewReader(os.Stdin)
	for {
		_, _ = fmt.Print("Enter command > ")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}
		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			_, _ = fmt.Println("Unrecognized command:", tokens[0])
		}
	}
}
