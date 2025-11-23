package main

import "fmt"

type MusicDownloader struct{}

func (m *MusicDownloader) Download(songUrl string) {

	fmt.Println("Downloading song: ", songUrl)
}

type MusicPlayer struct{}

func (p *MusicPlayer) Play(song string) {
	fmt.Println("Playing song: ", song)
}

type PlayListManager struct {
	playlist []string
}

func (pl *PlayListManager) AddToPlayList(song string) {
	pl.playlist = append(pl.playlist, song)
	fmt.Println("Added song to playlist : ", song)
}

type MusicFacade struct {
	downloader      *MusicDownloader
	player          *MusicPlayer
	playListManager *PlayListManager
}

func NewMusicFacade() *MusicFacade {
	return &MusicFacade{
		downloader:      &MusicDownloader{},
		player:          &MusicPlayer{},
		playListManager: &PlayListManager{},
	}
}

func (mf *MusicFacade) PlayMusic(song string) {
	mf.downloader.Download(song)
	mf.playListManager.AddToPlayList(song)
	mf.player.Play(song)

}

func main() {

	music := NewMusicFacade()
	music.PlayMusic("shadmehr...")

}
