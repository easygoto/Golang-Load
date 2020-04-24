package library

import (
	"errors"
	"fmt"
)

type MusicEntry struct {
	Id     string
	Name   string // 音乐名称
	Artist string // 歌手
	Source string // 资源
	Type   string // 音乐风格
	Genre  string // 音乐格式
}

type MusicManager struct {
	musicList []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

func (musicManager *MusicManager) Len() int {
	return len(musicManager.musicList)
}

func (musicManager *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(musicManager.musicList) {
		return nil, errors.New(`"Index out of range."`)
	}
	return &musicManager.musicList[index], nil
}

func (musicManager *MusicManager) Find(name string) *MusicEntry {
	if len(musicManager.musicList) == 0 {
		return nil
	}
	for _, m := range musicManager.musicList {
		if m.Name == name {
			return &m
		}
	}
	return nil
}

func (musicManager *MusicManager) Add(music *MusicEntry) {
	for _, item := range musicManager.musicList {
		if item.Name == music.Name {
			_, _ = fmt.Println("music already exists!", music)
			return
		}
	}
	musicManager.musicList = append(musicManager.musicList, *music)
	_, _ = fmt.Println("music add success!")
}

func (musicManager *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(musicManager.musicList) {
		return nil
	}
	removedMusic := &musicManager.musicList[index]

	if removedMusic != nil {
		_, _ = fmt.Println("remove music success!", removedMusic)
	} else {
		_, _ = fmt.Println("music is not exists!")
	}

	if len(musicManager.musicList) == 0 {
		musicManager.musicList = make([]MusicEntry, 0)
	} else if index == 0 {
		musicManager.musicList = musicManager.musicList[1:]
	} else if index == len(musicManager.musicList)-1 {
		musicManager.musicList = musicManager.musicList[:index-1]
	} else {
		musicManager.musicList = append(musicManager.musicList[:index-1], musicManager.musicList[index+1:]...)
	}
	return removedMusic
}

func (musicManager *MusicManager) RemoveByName(name string) *MusicEntry {
	if name == "" || len(musicManager.musicList) == 0 {
		return nil
	}

	for index, music := range musicManager.musicList {
		if name == music.Name {
			return musicManager.Remove(index)
		}
	}
	return nil
}
