package main

// Реализовать паттерн «адаптер» на любом примере.
// https://dev-gang.ru/article/shablon-adaptera-v-golang-5obr8sso32/
import "fmt"

type Mp3Player interface {
	playMp3(fileType, filename string)
}

type AudioPlayer interface {
	playAudio(filename string) string
	createLyrics(filename string) string
}

type FancyAudioPlayer struct{}

func (ap *FancyAudioPlayer) createLyrics(filename string) string {
	return fmt.Sprintf("That is text for yours %s", filename)
}
func (ap *FancyAudioPlayer) playAudio(filename string) string {
	return fmt.Sprintf("Playing cool %s", filename)
}

type Mp3PlayerAdapter struct {
	audioPlayer AudioPlayer
}

func (mp3P *Mp3PlayerAdapter) playMp3(fileType, filename string) {

	if fileType == "aac" {
		fmt.Printf("%s\n", mp3P.audioPlayer.playAudio(filename))
	} else if fileType == "lyrics" {
		fmt.Printf("%s\n", mp3P.audioPlayer.createLyrics(filename))
	} else {
		fmt.Printf("Ошибка!\n")
	}
}

type CoolPlayer struct {
	mp3PlayerAdapter Mp3Player
}

func (cp *CoolPlayer) playMp3(fileType, filename string) {
	if fileType == "mp3" {
		fmt.Printf("Playing Mp3 %s\n", filename)
	} else if fileType == "lyrics" || fileType == "aac" {
		cp.mp3PlayerAdapter = &Mp3PlayerAdapter{&FancyAudioPlayer{}}
		cp.mp3PlayerAdapter.playMp3(fileType, filename)
	}

}

func main() {
	cp := CoolPlayer{}

	cp.playMp3("mp3", "lalala")
	cp.playMp3("aac", "lalala")
	cp.playMp3("lyrics", "lalala")
}
