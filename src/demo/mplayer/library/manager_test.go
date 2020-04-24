package library

import (
	"reflect"
	"testing"
)

func TestOps(t *testing.T) {
	musicManager := NewMusicManager()
	if musicManager == nil {
		t.Error("NewMusicManager failed.")
		return
	}
	if musicManager.Len() != 0 {
		t.Error("NewMusicManager failed, not empty.")
		return
	}

	music0 := &MusicEntry{
		Id: "1", Name: "My Heart Will Go On", Artist: "Celion Dion", Source: "http://xxx", Type: "Pop", Genre: "MP3"}
	musicManager.Add(music0)
	if musicManager.Len() != 1 {
		t.Error("MusicManager.Add() failed.")
		return
	}

	music := musicManager.Find(music0.Name)
	if music == nil {
		t.Error("MusicManager.Find() failed.")
		return
	}

	if !reflect.DeepEqual(music, music0) {
		t.Error("MusicManager.Find() failed. Found item mismatch.")
		return
	}

	music, err := musicManager.Get(0)
	if music == nil {
		t.Error("MusicManager.Get() failed.", err)
		return
	}

	music = musicManager.Remove(0)
	if music == nil || musicManager.Len() != 0 {
		t.Error("MusicManager.Remove() failed.", err)
	}
}
